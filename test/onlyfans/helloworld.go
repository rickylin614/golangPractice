package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func main() {

	url := "https://onlyfans.com/api2/v2/users/81082076/posts/medias?limit=10&order=publish_date_desc&skip_users=all&format=infinite&counters=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7,ja;q=0.6,zh-CN;q=0.5,es;q=0.4")
	req.Header.Set("app-token", "33d57ade8c02dbc5a333db99ff9ae26a")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sign", "7865:30c67b6d916c219dd7125d01b9e67f04101e5d49:8af:6436c06b")
	req.Header.Set("user-id", "42987632")
	req.Header.Set("x-bc", "34858ba813853df791e981fb889bec868edab7c3")
	req.Header.Set("cookie", "auth_hash=06e377a837140d35b828b3bce5bc8195; csrf=SZ6JeorUf5f3b74ab20548899bcba17516b189b7; auth_id=42987632; cookiesAccepted=all; __stripe_mid=b594bb4e-a0fe-4f91-98a4-bc38c30e7cfa262046; fp=e957566f44bdf6f54bfca1ad422b5b95; sess=pqbijcs23d47lpqktpmftt0g9t; st=b3bc26ee3358f8ab6d5dc693b1926543720ae4f12c50301d78d9ffa8e01d0154; ref_src=")
	req.Header.Set("Referer", "https://onlyfans.com/meenfox/media")
	req.Header.Set("Referrer-Policy", "strict-origin-when-cross-origin")
	t := time.Now()
	tr := strconv.Itoa(int(t.UnixMilli()))
	req.Header.Set("time", tr)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		fmt.Printf("%s", body)
		return
	}

	m := make(map[string]any, 0)
	pictrue := make(map[string]any, 0)
	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	if v1s, ok := m["list"].([]map[string]any); ok {
		for _, v1 := range v1s {
			if v2s, ok := v1["media"].([]map[string]any); ok {
				for _, v2 := range v2s {
					if v3, ok := v2["full"].(string); ok {
						pictrue[v3] = true
					}
				}
			}
		}
	}

	// fmt.Println(res)
	// fmt.Println(string(body))
	fmt.Println(pictrue)

}
