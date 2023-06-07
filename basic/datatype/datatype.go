package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
)

func main() {
	intTest()
	arrayTest()
	sliceTest()
}

func intTest() {
	var num8_max int8 = int8(^uint8(0) >> 1)
	var num16_max int16 = int16(^uint16(0) >> 1)
	var num32_max int32 = int32(^uint32(0) >> 1)
	var num64_max int64 = int64(^uint(0) >> 1)
	var num_max int = int(^uint(0) >> 1)
	fmt.Printf("int8: %d to %d\n", ^num8_max, num8_max)
	fmt.Printf("int16: %d to %d\n", ^num16_max, num16_max)
	fmt.Printf("int32: %d to %d\n", ^num32_max, num32_max)
	fmt.Printf("int64: %d to %d\n", ^num64_max, num64_max)
	fmt.Printf("int: %d to %d\n", ^num_max, num_max)
	var unum8 uint8 = 0
	var unum16 uint16 = 0
	var unum32 uint32 = 0
	var unum64 uint64 = 0
	var nnum uint = 0
	fmt.Printf("uint8: %d to %d\n", unum8, ^unum8)
	fmt.Printf("uint16: %d to %d\n", unum16, ^unum16)
	fmt.Printf("uint32: %d to %d\n", unum32, ^unum32)
	fmt.Printf("uint64: %d to %d\n", unum64, ^unum64)
	fmt.Printf("uint: %d to %d\n", nnum, ^nnum)
}
func arrayTest() {
	// var "變數名稱" ["陣列長度"] "元素型態"
	var intArray [10]int = [10]int{1, 2, 3, 4, 5}        //未宣告值預設給0
	var stringArray [10]string = [10]string{"1", "2"}    //未宣告值預設給空字串
	var floatArray [10]float64 = [10]float64{1.00, 2.00} //未宣告值預設給0.000000
	intArray[2] = 999
	stringArray[9] = "test"
	/* 输出每个数组元素的值 */
	fmt.Println(intArray)
	fmt.Println(stringArray)
	fmt.Println(floatArray)
}

func sliceTest() {
	//make() 創建一個slice
	// var "變數名稱" [] "元素型態" 或使用 make函數
	var intSlice []int = make([]int, 2, 10) // 2 = 創建長度
	var intSliceN []int = []int{1, 2, 3}
	intSlice[1] = 9
	intSlice = append(intSlice, 55)                   //append 擴充slice使用
	intSlice = append(intSlice, 1, 2, 3, 4, 5, 6, 66) //append 擴充slice使用
	fmt.Print(intSlice)
	fmt.Print(intSliceN)
}

// func ToInt(str string) int {
// 	i, _ := strconv.Atoi(str)
// 	// _ = 拋棄符 用來處理不需要使用的變數
// 	return i
// }

// func swat( *a int, *b int) {
// 	a, b = b, a
// }

func ToString(arg interface{}, timeFormat ...string) string {
	if len(timeFormat) > 1 {
		log.SetFlags(log.Llongfile | log.LstdFlags)
		log.Println(fmt.Errorf("timeFormat's length should be one"))
	}
	var tmp = reflect.Indirect(reflect.ValueOf(arg)).Interface()
	switch v := tmp.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case string:
		return v
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case time.Time:
		if len(timeFormat) == 1 {
			return v.Format(timeFormat[0])
		}
		return v.Format("2006-01-02 15:04:05")
	// case jsoncrack.Time:
	// if len(timeFormat) == 1 {
	// return v.Time().Format(timeFormat[0])
	// }
	// return v.Time().Format("2006-01-02 15:04:05")
	case reflect.Value:
		return ToString(v.Interface(), timeFormat...)
	case fmt.Stringer:
		return v.String()
	default:
		return ""
	}
}
