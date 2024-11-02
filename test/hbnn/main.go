package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

const testCount = 100000
const oldAlgorithm = false

func init() {
	// mrand.Seed(time.Now().UnixNano())
}

func main() {
	Batch()
	// fmt.Println(RandMoney(1, 2))
}

func Batch() {

	// 總額 & 包數
	// for i := 5; i < 200; i = i + 5 {
	// for qty := 2; qty < 15; qty++ {
	// 	Settlement(4, qty)
	// }
	// }

	Settlement(1, 4)
}

func Settlement(total float64, qty int) {
	f := make(map[int]int, 0)

	// 總合
	for i := 0; i < testCount; i++ {
		// 分派
		t := total
		for i := 0; i < qty; i++ {
			money := RandMoney(t, qty-i)
			if qty-i != 1 {
				f[getPoint(money)]++
			}
			t -= money
		}
	}

	fmt.Printf("f %+v\n", f)
}

func RandMoney(total float64, qty int) float64 {
	if total == 0 {
		panic(nil)
	}
	if qty == 1 {
		return total
	}

	// 1. 减去 qty * 0.2 的金额
	t := int(total*100) - (qty * 20)
	if t == 0 { // 防呆
		t = 1
	}

	// 2. 计算上限值
	max := t*2/qty - 1
	if max < 10 { // 防呆
		max = 10
	}

	// 3. 随机预计牛数
	expectedBulls := rand.Intn(10)

	// 4. 随机分配数
	distribution := rand.Intn(max/10)*10 + 10

	// 5. 计算某数并加回分配数
	remainder := expectedBulls - ((distribution/100)+(distribution/10)%10)%10
	if remainder < 0 {
		remainder = 10 + remainder
	}
	if remainder < 0 || remainder > 10 {
		fmt.Println("異常1", remainder, distribution, expectedBulls)
	}
	if distribution == t {
		fmt.Println("異常2", remainder, distribution, expectedBulls, qty)
	}
	distribution += remainder
	if float64(distribution)/100 == 0 {
		fmt.Println("異常3", remainder, distribution, expectedBulls, qty)
	}
	result := float64(distribution) / 100
	pp := getPoint(result)
	if pp%10 != expectedBulls {
		fmt.Println("異常4", pp%10, result, remainder, distribution, expectedBulls, qty)
	}
	if result == 0 || result < 0.01 {
		fmt.Println("異常5", result, remainder, distribution, expectedBulls, qty)
	}

	// 6. 返回最终分配数
	return result
}

func getPoint(amount float64) int {
	num := strings.Replace(strconv.FormatFloat(round(amount, .5, 2), 'f', 2, 64), ".", "", -1)
	n := len(num)
	n1 := atoi(string(num[n-1]), -1)
	n2 := atoi(string(num[n-2]), -1)
	n3 := atoi(string(num[n-3]), -1)

	if n1 == -1 {
		return -1
	}

	if n2 == -1 {
		return -1
	}

	if n3 == -1 {
		return -1
	}

	nn := (n1 + n2 + n3) % 10
	if nn == 0 {
		return 10
	}

	return nn

}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)

	_div := math.Copysign(div, val)
	_roundOn := math.Copysign(roundOn, val)

	if _div >= _roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}

	newVal = round / pow
	return
}

func atoi(v string, d int) int {

	s := strings.Trim(v, "\"")

	if len(s) == 0 {
		return d
	} else {
		s = strings.TrimLeft(s, "0")
		if len(s) == 0 { //v == "0"
			return 0
		}
	}

	if i, err := strconv.Atoi(s); err != nil {
		return d
	} else {
		return i
	}
}
