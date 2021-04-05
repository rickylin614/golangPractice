package main

import (
	"fmt"
	"io/ioutil"
	"myGolang/ch16_elasticsearch/elasticeDemo"
	"net/http"
	"regexp"
	"strings"
	"time"
)

//從零開始建立分布式爬蟲 -- 不使用現成的爬蟲庫/框架

// get data from Internet and store it into datebase then create a front to show it

// go 的爬蟲框架
// henrylee2cn/pholcus
// colly
// gocrawl
// hu17889/go_spider

// 使用 ElasticSearch 作為數據存儲
//	GO語言標準模板庫實現http數據展示

var url string = "https://gnn.gamer.com.tw/index.php"
var indexPage [7]string = [7]string{"4", "1", "2", "5", "13", "11", "9"}

type myBody string

func (body myBody) Read(p []byte) (n int, err error) {
	return strings.NewReader(string(body)).Read(p)
}

func main() {
	// req, err := http.NewRequest(http.MethodGet, url, nil)
	// handleErr(err)

	urlWithParam := url + "?k=" + indexPage[0]
	fmt.Println(urlWithParam)
	resp, err := http.Get(urlWithParam)
	handleErr(err)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		handleErr(err)
		myRegExp(all)
	}
	time.Sleep(time.Second * 10)
}

var titleReg string = "<a.+gamer.com.tw.+a>"

type newsData struct {
	Sn    string
	Title string
	Url   string
}

func myRegExp(data []byte) {
	re := regexp.MustCompile(titleReg)
	match := re.FindAllString(string(data), -1)
	m := make(map[string]newsData)
	// endSli := make([]string, len(match))
	c := saveWorker() //做一個goroutine儲存訊息用 返回保存用channel
	for _, v := range match {
		if !strings.Contains(v, "繼續閱讀") && !strings.Contains(v, "img src=") && strings.Contains(v, "gnn.gamer.com.tw") {
			// endSli = append(endSli, v)
			composeNews(v, m, c)
			// fmt.Println(v)
		}
	}
	// fmt.Println(m)
	// fmt.Println(match)
}

// func storeFile(m map[string]newsData) {
// 	ioutil.WriteFile("list.txt", "", os.mode)
// 	for k, v := range m {

// 	}
// }

func composeNews(v string, m map[string]newsData, c chan newsData) {
	// fmt.Println(v)
	titleReg := regexp.MustCompile(">.*</a>")
	tm := titleReg.FindString(v)
	if len(tm) > 4 {
		tm = tm[1 : len(tm)-4]
	}
	urlReg := regexp.MustCompile("gnn.gamer.com.*sn=[0-9]{1,6}")
	match := urlReg.FindString(v)

	if len(match) > 6 {
		sn := match[len(match)-6:]
		data := newsData{
			sn,
			tm,
			match,
		}
		m[sn] = data
		c <- data
	}
}

func saveWorker() chan newsData {
	c := make(chan newsData, 1024) //創channel
	go func() {
		for {
			datas := <-c
			// fmt.Println(datas)
			elasticeDemo.StoreIntoElastic(datas, "gnn", "gamer") //保存資料
		}
	}()
	return c
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
