package main

import (
	"myGolang/ch16_elasticsearch/elasticeDemo"
)

//一般操作
//<server>: 9200 / index / type / id
// index -> database name
// type -> table name
// id -> unique key for object
//PUT / POST 新增修改資料
//GET取得資料

//<server>: 9200 / index / type / _search // 查詢所有資料

// type newsData struct {
// 	Sn    string
// 	Title string
// 	Url   string
// }

//Elasticsearch 7.x 后 所有_type已经不再使用 路径当中以_doc取代 type
//若要区分两种不同的db-table 则是使用组合字串 user-username / pet-petinfo

func main() {
	db := "gnn"
	// table := "gamer"

	// a := newsData{"aaa", "bbb", "ccc"}
	// elasticeDemo.StroeByElastic(a, db, table)
	elasticeDemo.SearchByEalstic("930rl3UB8i6m0td4X21t", db)
	// fmt.Println(str)
	// fmt.println
	// ReflectTest(a)
}

// func ReflectTest(obj interface{}) {
// 	fmt.Println(obj)

// 	//Value
// 	// v := reflect.ValueOf(obj)
// 	fmt.Println(structs.New(obj).Values()...)
// 	fmt.Println(structs.Names(obj))
// 	// fmt.Println(structs.Map(obj))

// 	out := structs.Map(obj)
// 	jsonStr, _ := json.Marshal(out)

// 	fmt.Println(out)
// 	fmt.Println(string(jsonStr))
// 	// fmt.Println(v.Elem().String())
// }
