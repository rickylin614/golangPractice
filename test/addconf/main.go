package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var (
	envs = []string{"dev"}
	// brands = []string{"3h", "bh", "c7", "c8", "cdd", "co", "hy", "ls", "lv", "sc", "ttmj", "tz", "xpj", "demo", "demo2", "ia"}
	// envs = []string{"dev", "uat", "stage", "prod"}
	// brands = []string{"3h", "bh", "c7", "c8", "cdd", "co", "hy", "ls", "lv", "sc", "ttmj", "tz", "xpj"}
	// brands = []string{"3h", "bh", "c7", "c8", "cdd", "co", "hy", "ls", "lv", "sc", "ttmj", "tz", "xpj", "demo", "demo2", "ia", "br001", "vn001", "in002", "template"}
	brands = []string{"c7"}
)

const (
	branch        = "feature/ricky/BE1-4365__返水自動生成2" // 暫存分支名稱
	commitMessage = "feat: BE1-4365__返水自動生成2"         // commit訊息
	MRmessage     = "返水自動生成2"                         // MR訊息
	productName   = "xunya-service"                   // 專案種類
	fileName      = "app.conf"                        // 修改的檔案名稱
	accessToken   = "glpat-42LtXjYn5G9RuE2ZzHjE"      // git accessToken
)

// 欲添加的參數, title: content
var AddConfig = []Config{
	// {"[mission_sys]", "request_processing_timeout=48"},
	// {"[mission_sys]", "status_update_fetch_range=24"},
	// {"[mission_sys]", "process_reload_fetch_range=18"},
	// {"[mission_sys]", "payout_request_topic=mission-payout-"},
	// {"[mission_sys]", "payout_result_topic=mission-payout-result"},
	// {"[mission_sys]", "payout_message_result_topic=mission-mail-result"},
	{"[autorebate]", "enable=true"},
	{"[autorebate]", "cron=0 30 15 * * *1"},
	{"[autorebate]", "rebate_id=dd3468d4-7278-46da-8449-09cb7c5106d9"},
	// {"[mission_sys]", "enable=true"},
}

var ProjectMapping map[string]string = map[string]string{
	"xunya-apis":    "280",
	"xunya-service": "281",
	"xunya-moapi":   "284",
	"central-proxy": "1732",
}

type Config struct {
	Title   string
	Content string
}

func main() {
	gitClone()
	defer removeProject()

	for _, v := range AddConfig {
		addLine(v.Title, v.Content)
	}

	for _, v := range envs {
		env := v
		gitAddFile(env)
		gitCommitPush(env)
		CreateMergeRequest(env)
	}

}

//	func gitClone() {
//		cmd := exec.Command("cmd", "/C,", "git_clone.bat", productName)
//		out, err := cmd.Output()
//		if err != nil {
//			fmt.Println("could not run command git_clone: ", err)
//		}
//		fmt.Println("Output: ", string(out))
//	}
func gitClone() {
	// 執行 git clone
	gitCloneCmd := exec.Command("git", "clone", "http://gitlab.paradise-soft.com.tw/configuration/"+productName+".git")
	gitCloneCmd.Stdout = os.Stdout
	gitCloneCmd.Stderr = os.Stderr
	if err := gitCloneCmd.Run(); err != nil {
		fmt.Printf("Error running git clone: %v\n", err)
		os.Exit(1)
	}

	// 切換到下載的目錄並執行 git pull
	os.Chdir(productName)
	gitPullCmd := exec.Command("git", "pull")
	gitPullCmd.Stdout = os.Stdout
	gitPullCmd.Stderr = os.Stderr
	if err := gitPullCmd.Run(); err != nil {
		fmt.Printf("Error running git pull: %v\n", err)
		os.Exit(1)
	}

	// 回到原來的目錄
	os.Chdir("..")
}

// func removeProject() {
// 	cmd := exec.Command("cmd", "/C,", "git_del.bat", productName)
// 	out, err := cmd.Output()
// 	if err != nil {
// 		fmt.Println("could not run command git_del: ", err)
// 	}
// 	fmt.Println("Output: ", string(out))
// }

func removeProject() {
	err := os.RemoveAll(productName)
	if err != nil {
		fmt.Printf("Error removing project folder: %v\n", err)
	} else {
		fmt.Println("Product folder has been deleted.")
	}
}

var configs map[string]map[string]string = make(map[string]map[string]string)

func addLine(targetTitle string, addString string) {
	for _, v := range envs {
		env := v
		for _, v2 := range brands {
			brand := v2
			filePath := "./" + productName + "/" + env + "/" + brand + "/" + fileName

			readFile, err := os.Open(filePath)

			if err != nil {
				fmt.Println(err)
			}
			fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)
			var fileLines []string

			nowTitle := ""
			for fileScanner.Scan() {
				// 取得單行訊息
				text := fileScanner.Text()
				text = strings.TrimSpace(text)

				// 分析是否為標題[]
				titleRegexp := regexp.MustCompile(`^\[.*\]$`)
				if titleRegexp.MatchString(text) {
					nowTitle = text
					continue
				}

				// 跳過註解
				if strings.HasPrefix(text, "#") {
					continue
				}

				// 將資料放入分類內
				content := strings.Split(text, "=")
				if len(content) == 2 {
					if _, ok := configs[nowTitle]; ok {
						configs[nowTitle][content[0]] = content[1]
					} else {
						configs[nowTitle] = map[string]string{
							content[0]: content[1],
						}
					}
				}

			}
			readFile.Close()
			// 重新讀檔
			readFile, err = os.Open(filePath)
			if err != nil {
				fmt.Println(err)
			}
			fileScanner = bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)
			nowTitle = ""
			for fileScanner.Scan() {
				// 取得單行訊息
				text := fileScanner.Text()
				text = strings.TrimSpace(text)

				// 預計添加的變數
				target := strings.Split(addString, "=")

				// 分析是否為標題[]
				titleRegexp := regexp.MustCompile(`^\[.*\]$`)
				if titleRegexp.MatchString(text) {
					nowTitle = text
					fileLines = append(fileLines, text)
					if len(target) == 2 && nowTitle == targetTitle {
						if _, ok := configs[nowTitle][target[0]]; !ok { // 表示沒有重複 需要新增
							fileLines = append(fileLines, addString)
						}
					}
					continue
				}

				// 分析內容 有重複做修改 沒重複正常新增
				contents := strings.Split(text, "=")
				if len(target) == 2 && len(contents) == 2 && target[0] == contents[0] && nowTitle == targetTitle { // 重複
					fileLines = append(fileLines, addString)
				} else {
					fileLines = append(fileLines, text)
				}
			}

			if _, ok := configs[targetTitle]; !ok {
				fileLines = append(fileLines, "")
				fileLines = append(fileLines, targetTitle)
				fileLines = append(fileLines, addString)
			}

			readFile.Close()
			os.Remove(filePath)
			writeFile, err := os.Create(filePath)
			datawriter := bufio.NewWriter(writeFile)

			for _, line := range fileLines {
				_, _ = datawriter.WriteString(line + "\n")
			}

			datawriter.Flush()
			writeFile.Close()
		}
	}
}

// func gitAddFile(env string) {
// 	for _, v2 := range brands {
// 		brand := v2
// 		cmd := exec.Command("cmd", "/C,", "git_add_file.bat", productName, brand, fileName, env)
// 		out, err := cmd.Output()
// 		if err != nil {
// 			// if there was any error, print it here
// 			fmt.Println("could not run command git_add_file.bat: ", err)
// 			fmt.Println("env: ", env, " brand: ", brand)
// 		}
// 		fmt.Println("Output: ", string(out))
// 	}
// }

func gitAddFile(env string) {
	for _, brand := range brands {
		filePath := "./" + productName + "/" + env + "/" + brand + "/" + fileName
		gitAddCmd := exec.Command("git", "add", filePath)
		gitAddCmd.Stdout = os.Stdout
		gitAddCmd.Stderr = os.Stderr
		if err := gitAddCmd.Run(); err != nil {
			fmt.Printf("Error running git add for %s: %v\n", filePath, err)
		} else {
			fmt.Printf("Added %s to git staging\n", filePath)
		}
	}
}

// func gitCommitPush(env string) {
// 	cmd := exec.Command("cmd", "/C,", "git_commit_push.bat", productName, "brand", commitMessage, branch+"_"+strings.ToUpper(env), env)
// 	out, err := cmd.Output()
// 	if err != nil {
// 		// if there was any error, print it here
// 		fmt.Println("could not run command \"git_commit_push.bat: ", err)
// 		fmt.Println("env: ", env)
// 	}
// 	fmt.Println("Output: ", string(out))
// }

func gitCommitPush(env string) {
	branchName := branch + "_" + strings.ToUpper(env)
	os.Chdir(productName)

	// Checkout master branch
	gitCheckoutMasterCmd := exec.Command("git", "checkout", "master")
	runCommand(gitCheckoutMasterCmd)

	// Pull the latest changes
	gitPullCmd := exec.Command("git", "pull")
	runCommand(gitPullCmd)

	// Create a new branch
	gitCreateBranchCmd := exec.Command("git", "checkout", "-b", branchName)
	runCommand(gitCreateBranchCmd)

	// Add changes
	gitAddCmd := exec.Command("git", "add", ".")
	runCommand(gitAddCmd)

	// Commit changes
	gitCommitCmd := exec.Command("git", "commit", "-m", commitMessage)
	runCommand(gitCommitCmd)

	// Push the changes
	gitPushCmd := exec.Command("git", "push", "origin", branchName)
	runCommand(gitPushCmd)

	os.Chdir("..")
}

func runCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}
}

// 尚未測試。可能有BUG
func gitCommitPush2(env string) {
	branchName := branch + "_" + strings.ToUpper(env)

	// Checkout the master branch
	gitCheckoutMasterCmd := exec.Command("git", "checkout", "master")
	gitCheckoutMasterCmd.Stdout = os.Stdout
	gitCheckoutMasterCmd.Stderr = os.Stderr

	if err := gitCheckoutMasterCmd.Run(); err != nil {
		fmt.Printf("Error checking out master branch: %v\n", err)
		os.Exit(1)
	}

	// Pull the latest changes
	gitPullCmd := exec.Command("git", "pull")
	gitPullCmd.Stdout = os.Stdout
	gitPullCmd.Stderr = os.Stderr

	if err := gitPullCmd.Run(); err != nil {
		fmt.Printf("Error pulling latest changes: %v\n", err)
		os.Exit(1)
	}

	// Create a new branch
	gitCreateBranchCmd := exec.Command("git", "checkout", "-b", branchName)
	gitCreateBranchCmd.Stdout = os.Stdout
	gitCreateBranchCmd.Stderr = os.Stderr

	if err := gitCreateBranchCmd.Run(); err != nil {
		fmt.Printf("Error creating branch %s: %v\n", branchName, err)
		os.Exit(1)
	}

	// Commit changes with the provided commit message
	gitCommitCmd := exec.Command("git", "commit", "-m", commitMessage)
	gitCommitCmd.Stdout = os.Stdout
	gitCommitCmd.Stderr = os.Stderr

	if err := gitCommitCmd.Run(); err != nil {
		fmt.Printf("Error committing changes: %v\n", err)
		os.Exit(1)
	}

	// Push Commit
	gitPushCmd := exec.Command("git", "push", "origin", branchName)
	gitPushCmd.Stdout = os.Stdout
	gitPushCmd.Stderr = os.Stderr

	if err := gitPushCmd.Run(); err != nil {
		fmt.Printf("Error git push: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully performed the Git operations for %s in %s branch\n", productName, branchName)
}

type MergeRequest struct {
	SourceBranch       string `json:"source_branch"`
	TargetBranch       string `json:"target_branch"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	RemoveSourceBranch bool   `json:"remove_source_branch"`
}

// generatered by chatgpt
func CreateMergeRequest(env string) {
	gitlabURL := "https://gitlab.paradise-soft.com.tw"
	projectID := ""
	if v, ok := ProjectMapping[productName]; ok {
		projectID = v
	} else {
		fmt.Println("查無專案ID，取消創立Merge Request")
		return
	}

	sourceBranch := branch + "_" + strings.ToUpper(env)
	targetBranch := "master"
	title := MRmessage + "_" + strings.ToUpper(env)
	description := MRmessage + "_" + strings.ToUpper(env)
	removeSourceBranch := true

	mr := MergeRequest{
		SourceBranch:       sourceBranch,
		TargetBranch:       targetBranch,
		Title:              title,
		Description:        description,
		RemoveSourceBranch: removeSourceBranch,
	}

	data, err := json.Marshal(mr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	apiURL := fmt.Sprintf("%s/api/v4/projects/%s/merge_requests", gitlabURL, projectID)
	values := url.Values{}
	values.Set("access_token", accessToken)

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = values.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	mrResp := &MergeRequestResp{}

	err = json.Unmarshal(respBody, mrResp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", string(mrResp.WebURL))
}

type MergeRequestResp struct {
	ID                        int         `json:"id"`
	IID                       int         `json:"iid"`
	ProjectID                 int         `json:"project_id"`
	Title                     string      `json:"title"`
	Description               string      `json:"description"`
	State                     string      `json:"state"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt                 time.Time   `json:"updated_at"`
	MergedBy                  *User       `json:"merged_by"`
	MergeUser                 *User       `json:"merge_user"`
	MergedAt                  time.Time   `json:"merged_at"`
	ClosedBy                  *User       `json:"closed_by"`
	ClosedAt                  time.Time   `json:"closed_at"`
	TargetBranch              string      `json:"target_branch"`
	SourceBranch              string      `json:"source_branch"`
	UserNotesCount            int         `json:"user_notes_count"`
	Upvotes                   int         `json:"upvotes"`
	Downvotes                 int         `json:"downvotes"`
	Author                    User        `json:"author"`
	Assignees                 []User      `json:"assignees"`
	Assignee                  *User       `json:"assignee"`
	Reviewers                 []User      `json:"reviewers"`
	SourceProjectID           int         `json:"source_project_id"`
	TargetProjectID           int         `json:"target_project_id"`
	Labels                    []string    `json:"labels"`
	Draft                     bool        `json:"draft"`
	WorkInProgress            bool        `json:"work_in_progress"`
	Milestone                 interface{} `json:"milestone"`
	MergeWhenPipelineSucceeds bool        `json:"merge_when_pipeline_succeeds"`
	MergeStatus               string      `json:"merge_status"`
	SHA                       string      `json:"sha"`
	MergeCommitSHA            string      `json:"merge_commit_sha"`
	SquashCommitSHA           string      `json:"squash_commit_sha"`
	DiscussionLocked          interface{} `json:"discussion_locked"`
	ShouldRemoveSourceBranch  interface{} `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch   bool        `json:"force_remove_source_branch"`
	Reference                 string      `json:"reference"`
	References                struct {
		Short    string `json:"short"`
		Relative string `json:"relative"`
		Full     string `json:"full"`
	} `json:"references"`
	WebURL    string `json:"web_url"` // MR的網址
	TimeStats struct {
		TimeEstimate        int         `json:"time_estimate"`
		TotalTimeSpent      int         `json:"total_time_spent"`
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
	} `json:"time_stats"`
	Squash               bool `json:"squash"`
	TaskCompletionStatus struct {
		Count          int `json:"count"`
		CompletedCount int `json:"completed_count"`
	} `json:"task_completion_status"`
	HasConflicts                bool        `json:"has_conflicts"`
	BlockingDiscussionsResolved bool        `json:"blocking_discussions_resolved"`
	Subscribed                  bool        `json:"subscribed"`
	ChangesCount                interface{} `json:"changes_count"`
	LatestBuildStartedAt        interface{} `json:"latest_build_started_at"`
	LatestBuildFinishedAt       interface{} `json:"latest_build_finished_at"`
	FirstDeployedToProductionAt interface{} `json:"first_deployed_to_production_at"`
	Pipeline                    interface{} `json:"pipeline"`
	HeadPipeline                interface{} `json:"head_pipeline"`
	DiffRefs                    struct {
		BaseSHA  interface{} `json:"base_sha"`
		HeadSHA  interface{} `json:"head_sha"`
		StartSHA string      `json:"start_sha"`
	} `json:"diff_refs"`
	MergeError interface{} `json:"merge_error"`
	User       struct {
		CanMerge bool `json:"can_merge"`
	} `json:"user"`
	Message interface{} // 錯誤訊息
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}
