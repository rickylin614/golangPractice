package main

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"sort"

	"math/big"
	mrand "math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// n := 0
	max := 10000
	// qty := 6
	amount := 1
	r := NewRandomAmount()
	for i := 8; i < 20; i++ {
		submain(r, max, i, amount)
	}
}

func submain(r *RandomAmount, max, qty, amount int) {
	m := make(map[int]int, 0)

	for i := 10; i < max+10; i++ {
		var total float64 = float64(amount)

		for j := 0; j < qty; j++ {
			i1 := r.Next(strconv.Itoa(i)+"_"+strconv.Itoa(qty), 1*time.Second, total, qty, "1", false)
			num := int(i1*1000) % 100 / 10
			m[num]++
			total -= i1
		}
	}

	globaldigis := make(map[int]int, 0)
	for i := 0; i < 10; i++ {
		for j := 0; j < qty; j++ {
			globaldigis[i] = m[i]
		}
	}

	total := 0
	for _, v := range globaldigis {
		total += v
	}
	for i := 0; i < 10; i++ {
		f := float64(globaldigis[i]) / float64(total) * 100
		fmt.Printf("%d,%.02f,", globaldigis[i], f)
	}
	fmt.Println("")
}

// func PrintMapData(m map[int]int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf(",")
// 	}
// }

func Ftoa(v float64) string {
	return strconv.FormatFloat(Round64(v, .5, 2), 'f', 2, 64)
}

func Round64(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// var randLogger = log.NewLogger("hb.rand.money")

func calculationByHbsl(maxAmount int, maxRate int) int {
	tempMaxAmount := (maxAmount * maxRate)

	roundAmount, _ := rand.Int(rand.Reader, big.NewInt(int64(tempMaxAmount)))

	return int(roundAmount.Int64()) / 100
}

func RandMoney(total float64, qty int) float64 {
	//return float64(RandMoneyCent(int(total*100), qty)) / 100

	if qty == 1 {
		return total
	}

	max := int(total * 100)

	n := mrand.Intn(100) * max / qty * 2

	if n == 0 {
		return 0.01
	}

	return float64(n) / 10000
}

func RandMoneyCent(total int, qty int) int {
	defer func() {
		if r := recover(); r != nil {
			// randLogger.Errorln(r, "total=", total, " qty=", qty)
		}
	}()

	if total < 0 {
		// randLogger.Errorln(total, " rand ", qty)
		return 0
	}

	if qty == 1 {
		//先四舍五入再转回float64，防止精度问题
		// result, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", total), 64)
		// return result
		return total
	}

	max := total - 1*(qty-1)

	roundSilce := make([]int, 0, 10)
	for i := 0; i <= 10; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(99))

		randRate := int(n.Int64())
		// if randRate > 90 {
		//     randRate = 40
		// }

		if randRate == 0 {
			randRate = 1
		}

		roundSilce = append(roundSilce, calculationByHbsl(max, randRate))
	}

	roundIndex, _ := rand.Int(rand.Reader, big.NewInt(10))
	//i, _ := strconv.Atoi(fmt.Sprintf("%v", roundIndex))
	val := roundSilce[int(roundIndex.Int64())]

	if val == 0 {
		return 1
	}

	return val
}

//RandomAmount 红包随机金额
type RandomAmount struct {
	sync.RWMutex
	timeout chan *randmonValues
	values  map[string]*randmonValues
}

//NewRandomAmount 创建红包随机金额生成器
func NewRandomAmount() *RandomAmount {
	ra := &RandomAmount{}
	ra.timeout = make(chan *randmonValues, 10000)
	ra.values = make(map[string]*randmonValues)

	return ra
}

//RandmonValues 预分配的红包金额
type randmonValues struct {
	cancel    context.CancelFunc
	id        string
	addedTime time.Time
	values    []float64
}

func (ra *RandomAmount) waitToClear(ctx context.Context, id string, timeout time.Duration) {

	if timeout <= 0 {
		timeout = 1 * time.Minute
	}
	t := time.NewTicker(timeout)
	defer t.Stop()

	select {
	case <-t.C:
		ra.Lock()
		defer ra.Unlock()
		delete(ra.values, id)

		return

	case <-ctx.Done():
		return
	}

}

//Next 取出下一个金额
func (ra *RandomAmount) Next(id string, timeout time.Duration, amount float64, qty int, hitNumber string, antiHit bool) float64 {
	ra.Lock()
	defer ra.Unlock()

	values, ok := ra.values[id]
	if !ok {
		values = &randmonValues{}
		values.addedTime = time.Now()
		values.values = ra.splitAmounts(amount, qty, hitNumber, antiHit)
		ctx, cancel := context.WithCancel(context.Background())

		values.cancel = cancel

		go ra.waitToClear(ctx, id, timeout) // 時間到了清空ID柱列數據

		ra.values[id] = values
	}

	if qty == 1 {
		// value := values.values[0]
		values.cancel()
		delete(ra.values, id)
		return amount
	}

	l := len(values.values)

	if l == 0 {
		return 0
	}

	i, _ := rand.Int(rand.Reader, big.NewInt(int64(l)))

	next := int(i.Int64())

	value := values.values[next]

	values.values = append(values.values[:next], values.values[next+1:l]...)

	return value
}

func (ra *RandomAmount) splitAmounts(amount float64, qty int, hitNumber string, antiHit bool) []float64 {

	if qty == 1 {
		return []float64{amount}
	}

	var total int
	var digits []int64
	panicCount := 0
	for {
		total = int(amount * 100)

		// 新流程 若金額小於0.1*人數 可能有問題
		// 原理: 小數點二位先全隨機(1~10) 原本的分配模式只隨機小數點一位 最後兩者陣列相加
		digits = make([]int64, qty)
		digitsSum := 0
		for i := 0; i < qty-1; i++ {
			n, _ := rand.Int(rand.Reader, big.NewInt(10))
			n.Add(n, big.NewInt(1))
			digits[i] = n.Int64()
			digitsSum += int(n.Int64())
		}
		// 最後一組 計算餘數 小數二位的總數為10的倍數
		remainder := digitsSum % 10
		digits[qty-1] = 10 - int64(remainder)
		digitsSum += 10 - remainder

		total -= digitsSum
		total /= 10
		if total >= 0 {
			break // 防止亂數使total小於零的防呆
		}
		if panicCount > 10 { // 若超過10次total皆小於0則判斷為錯誤(理論上一次都不應該發生)
			panic(errors.New("error"))
			// TODO panic or return empty slice?
		}
		panicCount++
	}
	points := make([]int, qty-1, qty-1)
	max := big.NewInt(int64(total))

	for i := 0; i < qty-1; i++ {
		n := big.NewInt(0)
		if max.Cmp(n) != 0 {
			n, _ = rand.Int(rand.Reader, max)
		}
		points[i] = int(n.Int64())
	}

	sort.Ints(points)
	values := make([]int, qty, qty)
	var current int = 0
	var next int
	for i := 0; i < qty-1; i++ {
		next = points[i]
		values[i] = next - current
		current = next
	}
	values[qty-1] = total - current

	// 組合
	for i := range values {
		values[i] *= 10
		values[i] += int(digits[i])
	}

	if antiHit {
		adjust := 0
		for {
			for k, v := range values {
				if adjust > 0 {
					v = v + adjust
					adjust = 0
					values[k] = v
				} else if adjust < 0 {
					if (v + adjust) > 0 {
						v = v + adjust
						adjust = 0
						values[k] = v
					}
				}

				if ra.isHit(v, hitNumber) {
					if v > 1 {
						values[k] = v - 1
						//fmt.Println(v, "->", values[k])
						adjust++
					} else {
						values[k] = v + 1
						//fmt.Println(v, "->", values[k])
						adjust--
					}
				}
			}
			if adjust == 0 {
				break
			}
		}
	}

	floatValues := make([]float64, len(values), len(values))

	for k, v := range values {
		floatValues[k] = float64(v) / 100
	}

	return floatValues
}

func (ra *RandomAmount) isHit(value int, hitNumber string) bool {

	s := strconv.Itoa(value)

	return string(s[len(s)-1]) == hitNumber
}
