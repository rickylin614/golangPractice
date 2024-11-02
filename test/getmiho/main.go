package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	baseURL := "https://www.dto.jp/osaka/schedule-list?page="
	keyword := "みほ"

	page := 1
	found := false

	for !found {
		url := fmt.Sprintf("%s%d", baseURL, page)
		fmt.Printf("Fetching page: %s\n", url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching page: %v\n", err)
			break
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Printf("Received non-200 response: %d\n", resp.StatusCode)
			break
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			fmt.Printf("Error parsing page: %v\n", err)
			break
		}

		// Find the content that might contain the keyword
		doc.Find("span").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			if strings.Contains(text, keyword) {
				fmt.Printf("Found keyword on page %d\n", page)
				found = true
			}
		})

		time.Sleep(time.Microsecond * 100)

		if found {
			break
		}

		page++
	}

	if !found {
		fmt.Println("Keyword not found in any page.")
	}
}
