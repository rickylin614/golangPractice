package main

import "fmt"

func main() {
	s := []string{"1", "2", "3"}
	b := GetSlice()
	b = append(b, s...)

	fmt.Println(b)
}

func GetSlice() []string {
	return nil
}
