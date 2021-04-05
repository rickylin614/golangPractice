package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	euler()
	triangle()
}

//bool true,false
//int int8 int16 int32 int64
//uint uint8 uint16 uint32 uint64
//byte rune (rune為char類型的進階版 4byte的資料結構)
//float32,float64,complex64,complex128

func euler() {
	// i = root(-1) 根號-1
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c)) // 3^2 + 4^2 root = 5
	fmt.Printf("%.3f\n",
		cmplx.Pow(math.E, 1i*math.Pi)+1) // 無限接近0

}

func triangle() {
	var a, b int = 3, 4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
