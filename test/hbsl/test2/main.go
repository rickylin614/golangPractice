package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	d, _ := rand.Int(rand.Reader, big.NewInt(0))
	fmt.Println(d)
	if ok, err := okDemo(); err != nil {
		fmt.Println("clouds")
		return
	} else if ok {
		fmt.Println("456")
		return
	}

}

func okDemo() (bool, error) {
	return true, nil
}
