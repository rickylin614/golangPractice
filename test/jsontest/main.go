package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type SS struct {
	A string `json:"a,omitempty"`
	B string `json:"b,omitempty"`
	C string `json:"c,omitempty"`
}

type SSS struct {
	A int `json:"a,omitempty"`
	B int `json:"b,omitempty"`
	C int `json:"c,omitempty"`
}

type TT struct {
	T *time.Time `json:"a,omitempty" time_format:"2016"`
}

func main() {
	jsonStr := `{"a":1,"b":1.0,"c":1.00}`
	m := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println(err)
		return
	}

	a, _ := m["a"].(float64)
	b, _ := m["b"].(float64)
	c, _ := m["c"].(float64)
	fmt.Println(a, b, c)
	fmt.Println(int(a) == 1)
	// -----------------------------------------------------------
	ss := SS{}
	err = json.Unmarshal([]byte(jsonStr), &ss)
	if err != nil {
		fmt.Println(err)
		return
	}
	//--------------------------------------------------------------

	fmt.Println(m)

	t := time.Now()
	TT := TT{T: &t}
	tStr, _ := json.Marshal(TT)
	fmt.Println(string(tStr))

}
