package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) //int nil值 0 不存在的key給0
	var m3 map[string]int
	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(m2["name"]) //不存在的key给零值
	fmt.Println(m3["name"]) //不存在的key给零值

	//traversing map
	for key, value := range m { //可 for key or for _ , value
		fmt.Printf("key : %s , value : %s\n", key, value)
	}

	//getting map
	fmt.Println(m["name"])    //不存在的key给零值
	value, exist := m["name"] //预设回传值 值 存在
	fmt.Println(value, exist) // ccmouse , true

	//deleting value
	delete(m, "name")
	value2, exist2 := m["name"]
	fmt.Println(value2, exist2) // _ , false

	//len 取得数量
	fmt.Println(len(m)) //3

	//添加新數據
	m["name2"] = "cccmouse"
	fmt.Println(len(m)) //4
}
