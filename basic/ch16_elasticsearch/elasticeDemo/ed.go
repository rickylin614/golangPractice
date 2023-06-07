package elasticeDemo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fatih/structs"
	"github.com/olivere/elastic/v7"
)

type newsData struct {
	Sn    string
	Title string
	Url   string
}

var baseDomain string = "http://localhost:9200/"

func StroeByElastic(obj interface{}, db string) {
	cli, err := elastic.NewClient( //default connect to 127.0.0.1:9200
		//mest turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	iresp, err := cli.Index().Index(db).BodyJson(obj).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(iresp)
}

func SearchByEalstic(id string, db string) {
	cli, err := elastic.NewClient( //default connect to 127.0.0.1:9200
		//mest turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	iresp, err := cli.Get().Index(db).Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", iresp.Source)
	var actual newsData
	err = json.Unmarshal([]byte(iresp.Source), &actual)
	if err != nil {
		panic(err)
	}
	fmt.Printf("??? %v\n", actual)
	// return string(iresp.Source)
}

func DeleteByElastic(id string, db string) {
	cli, err := elastic.NewClient(
		//mest turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	iresp, err := cli.Delete().Id(id).Index(db).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(iresp)
}

func StoreIntoElastic(obj interface{}, db string, table string) {
	//設定域名請求資訊以及傳送內容
	urlDomain := baseDomain + db + "/" + table
	requestBody := convertDataToJson(obj)

	//設定傳送請求
	req, err := http.NewRequest(http.MethodPost, urlDomain, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	//送出請求並接收回傳值
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}

func convertDataToJson(s interface{}) []byte {
	m := structs.Map(s)
	jsonStr, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return jsonStr
}
