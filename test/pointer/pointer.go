package pointer

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
)

// 使用 reflect.ValueOf 和 v.Pointer() 取得變量的記憶體地址
func GetAddrReflect(x interface{}) (int, error) {
	v := reflect.ValueOf(x)
	if v.CanAddr() {
		return int(v.Pointer()), nil
	}
	return 0, errors.New("x is not a pointer")
}

// 取得反射取得記憶體第二種類
func GetAddrReflect2(x interface{}) (int, error) {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Ptr {
		return 0, errors.New("x is not a pointer")
	}
	return int(v.Pointer()), nil
}

// 使用 fmt.Sprintf("%p", x) 取得變量的記憶體地址
func GetAddrFmt(x interface{}) string {
	return fmt.Sprintf("%p", x)
}

// 生成一個包含 n 個元素的整數切片，每個元素的值在 0 到 999 之間
func GenerateIntSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(1000)
	}
	return s
}
