package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// http client
func main() {
	req, _ := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Mobile Safari/537.36")
	// resp, err := http.Get("https://www.imooc.com")
	// resp, err := http.DefaultClient.Do(req)
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
		// Timeout: time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}

//其他標準庫

//bufio log encoding/json regexp time strings math rand
