package main

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

const USER_NAME = "ricky_lin"
const PWD = "frog*0517"
const EMPTY_CONTENT = ""
const PROJECT = "BE1"
const ISSUE_TYPE = "Task"
const LABEL = "BEPFT_P_XUNYA"
const COMPONENT = "xunya"

var releatedIssue = "PFREQ-5026"

func main() {
	GeneratorIssue()
}

func GeneratorIssue() {
	tp := jira.BasicAuthTransport{
		Username: USER_NAME,
		Password: PWD,
	}

	client, err := jira.NewClient(tp.Client(), "https://jira.paradise-soft.com.tw/")
	if err != nil {
		fmt.Println("jira new client error", err)
		return
	}

	// 取得關聯單的資訊
	pmIssue, _, err := client.Issue.Get(releatedIssue, nil)
	// client.Issue.Delete("BE1-3334")
	if err != nil {
		fmt.Printf("jira get issue %+v client error: %s\n", pmIssue, err)
		return
	}

	// 取得創建單的資訊
	// createIssue, _, err := client.Issue.Get("BE1-3338", nil)
	createIssue, err := CreateNewIssue(client, pmIssue.Fields.Summary)
	if err != nil {
		fmt.Printf("jira create issue %+v client error: %s\n", pmIssue, err)
		return
	}

	err = UpdateIssue(client, createIssue.Key, pmIssue.Key)
	if err != nil {
		fmt.Printf("jira update issue %+v client error: %s\n", pmIssue, err)
		return
	}

}

// 修改Issue
func UpdateIssue(client *jira.Client, createdKey, ReleatdKey string) error {
	_, err := client.Issue.AddLink(&jira.IssueLink{
		Type: jira.IssueLinkType{
			Name: "Relates",
		},
		InwardIssue:  &jira.Issue{Key: createdKey},
		OutwardIssue: &jira.Issue{Key: ReleatdKey},
	})
	if err != nil {
		fmt.Println("jira add link err:", err)
		return err
	}

	// oldIssue, _, err := client.Issue.Get(createdKey, nil)
	// // oldIssue.Field

	// _, err = client.Issue.UpdateIssue(createdKey, map[string]any{
	// 	"customfield_10102":    map[string]any{"add": "BE1-191"},
	// 	"fieldsToForcePresent": "",
	// })
	// _, _, err = client.Issue.Update(oldIssue)

	// if err != nil {
	// 	fmt.Println("jira update issue err:", err)
	// 	return err
	// }

	return err
}

// 創建Issue
func CreateNewIssue(client *jira.Client, summary string) (*jira.Issue, error) {
	issue := jira.Issue{
		Fields: &jira.IssueFields{
			Assignee: &jira.User{
				Name: USER_NAME,
			},
			Type: jira.IssueType{
				Name: ISSUE_TYPE,
			},
			Project: jira.Project{
				Key: PROJECT,
			},
			Labels: []string{LABEL},
			Components: []*jira.Component{
				{Name: "xunya"},
			},
			Summary: summary,
		},
	}

	i, _, err := client.Issue.Create(&issue)
	if err != nil {
		fmt.Println("jira create issue err:", err)
		return nil, err
	}
	return i, nil
}
