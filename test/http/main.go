package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.DefaultClient.Timeout = time.Nanosecond
	resp, err := http.NewRequest("Get", `https://www.google.com`, nil)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf(`%s %s %s`, resp.Body, resp.Response.Status, err)
	}

}
