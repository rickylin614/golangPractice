package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	QQ     string `json:"qq"`
	Wechat string `json:"wechat"`
}

type Profile2 struct {
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	QQ       string `json:"qq"`
	Wechat   string `json:"wechat"`
	Telegram string `json:"telegram"`
}

func main() {
	data := Profile{
		Email:  "123",
		Mobile: "456",
		QQ:     "789",
		Wechat: "000",
	}

	buf, err := json.Marshal(&data)

	fmt.Println(string(buf), err)

	data2 := &Profile2{}

	err = json.Unmarshal(buf, data2)

	fmt.Println(data2, data2.Telegram, err)
}
