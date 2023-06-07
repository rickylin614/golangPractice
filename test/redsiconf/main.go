package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Redisconf struct {
	Auth           string `json:"auth"`
	Host           string `json:"host"`
	Name           string `json:"name"`
	Port           int    `json:"port,omitempty"`
	SshPort        int    `json:"ssh_port,omitempty"`
	TimeoutConnect int    `json:"timeout_connect,omitempty"`
	TimeoutExecute int    `json:"timeout_execute,omitempty"`
}

func main() {
	file, err := os.Open("./conf.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	repeat := make(map[string]bool, 0)
	result := make([]Redisconf, 0)

	for scan.Scan() {
		// 1. 取品牌
		brand := strings.ReplaceAll(scan.Text(), "IP(", "")
		brand = strings.ReplaceAll(brand, ")", "")
		if _, exist := repeat[brand]; exist {
			brand += " 2"
		} else {
			repeat[brand] = true
		}
		// 2. 取host / port
		if !scan.Scan() {
			break
		}
		addr := strings.Split(scan.Text(), ":")
		host := addr[0]
		port, _ := strconv.Atoi(addr[1])
		// 3. 取auth
		if !scan.Scan() {
			break
		}
		auth := scan.Text()
		// 組合
		result = append(result, Redisconf{
			Auth:           auth,
			Host:           host,
			Name:           "stage " + brand,
			Port:           port,
			SshPort:        22,
			TimeoutConnect: 60000,
			TimeoutExecute: 60000,
		})
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(result)
	fmt.Println(string(b))
}
