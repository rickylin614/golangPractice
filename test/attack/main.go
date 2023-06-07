package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urlTemplate := "http://xunya-apis.dev-cdd.svc.cluster.local:8000/v1/member/lockmember/%d"
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for j := 0; j < 1; j++ {
		for i := 1; i <= 10000; i++ {
			go func(i int) {
				url := fmt.Sprintf(urlTemplate, i+100000)
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Failed to send request: %s\n", err.Error())
				} else {
					fmt.Printf("Response from %s: %s\n", url, resp.Status)
					resp.Body.Close()
				}
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
}

// go tool pprof -http=:8080 -inuse_space http://xunya-apis.dev-cdd.svc.cluster.local:8000/debug/pprof/heap
// go tool pprof -http=:8081 -inuse_space http://xunya-apis.dev-cdd.svc.cluster.local:8000/debug/pprof/heap
// go tool pprof -http=:8082 -inuse_space http://xunya-apis.dev-cdd.svc.cluster.local:8000/debug/pprof/heap

// go tool pprof -http=:8080 -seconds 3 http://xunya-apis.dev-cdd.svc.cluster.local:8000/debug/pprof/profile

// go tool pprof -http=:8080 -inuse_space http://xunya-service.dev-lv.svc.cluster.local:8001/debug/pprof/heap
