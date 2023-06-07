package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Demo struct {
	A string     `json:"a,omitempty"`
	B string     `json:"b,omitempty"`
	C int        `json:"c,omitempty"`
	D *time.Time `json:"d,omitempty"`
	E float64    `json:"e,omitempty"`
	F int        `json:"f,omitempty"`
}

func (d Demo) CompareColumns() (cols []string) {
	return []string{"a", "b", "d", "e"}
}

type StatusTypeEnum int

const (
	StatusTypeEnum_Failed StatusTypeEnum = iota
	StatusTypeEnum_Success
	StatusTypeEnum_Process
)

func (s StatusTypeEnum) Dictionary() map[StatusTypeEnum]string {
	return map[StatusTypeEnum]string{
		StatusTypeEnum_Failed:  "失敗",
		StatusTypeEnum_Success: "成功",
		StatusTypeEnum_Process: "處理中",
	}
}

func (d Demo) ColumnsNames() map[string]map[int]string {
	m := map[string]map[int]string{
		"c": {
			0: "禁用",
			1: "啟用",
		},
		"f": {
			0: "未鎖定",
			1: "鎖定",
		},
	}
	return m
}

func (d Demo) ModuleName() map[string]string {
	// insert note, update note, delete note
	return map[string]string{}
}

type IDemo interface {
	CompareColumns() (cols []string)
	ColumnsNames() map[string]map[int]string
	ModuleName() map[string]string
}

func main() {
	t := time.Now()
	t2 := time.Now().Add(time.Hour)

	var ADemo interface{} = Demo{A: "A", B: "B", C: 0, D: &t, E: 1.001}
	var BDemo interface{} = Demo{A: "AA", B: "B", C: 10, D: &t2, E: 1.002}
	var v1, v2 IDemo

	v1, _ = ADemo.(IDemo)
	v2, _ = BDemo.(IDemo)

	ToMap(v1)

	// update
	if v1 != nil && v2 != nil {
		cols := v1.CompareColumns()
		fmt.Println(v1)
		fmt.Println(v2)
		fmt.Println(cols)
		m1 := StructToMap(v1)
		m2 := StructToMap(v2)
		for _, key := range cols {
			if m1[key] != m2[key] {
				switch m1[key].(type) {
				case float64:
					fmt.Printf("Diff Column: %v Before: %.0f , After: %.0f\n", key, m1[key], m2[key])
				case int:
					// TODO
				case string:
					fmt.Printf("Diff Column: %v Before: %v , After: %v\n", key, m1[key], m2[key])
				case time.Time:
					// TODO timeformat
					fmt.Printf("Diff Column: %v Before: %v , After: %v\n", key, m1[key], m2[key])
				default:
					fmt.Printf("Diff Column: %v Before: %v , After: %v\n", key, m1[key], m2[key])
				}

			}
		}
	} else if v1 != nil && v2 == nil {
		// TODO delete
	} else if v1 == nil && v2 != nil {
		// TODO insert
	}
}

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{}, 0)
	m, _ := json.Marshal(data)
	_ = json.Unmarshal(m, &result)
	return result
}

func ToMap(in interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		fmt.Printf("%+v\n", fi)
		if tagValue := fi.Tag.Get("json"); tagValue != "" {
			fmt.Println(tagValue)
		}
		// if tagValue := fi.Tag.Get(tagName); tagValue != "" {
		// 	out[tagValue] = v.Field(i).Interface()
		// }
	}
	return out, nil
}
