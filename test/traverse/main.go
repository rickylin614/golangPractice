package main

import "fmt"

// Global slice and map
var s1 = []string{}
var s2 = []string{}
var m1 = make(map[string]bool)
var m2 = make(map[string]bool)

// Function to add a string to a slice
func AddToSliceUnique(val string) {
	for _, v := range s1 {
		if v == val {
			return
		}
	}
	for _, v := range s2 {
		if v == val {
			return
		}
	}
	s1 = append(s1, val)
	return
}

// Function to add a string to a map
func AddToMapUnique(val string) {
	if _, ok := m1[val]; ok {
		return
	}
	if _, ok := m2[val]; ok {
		return
	}
	m1[val] = true
	return
}

func main() {
	strings := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	s := []string{}
	for _, v := range strings {
		AddToSliceUnique(v)
	}
	fmt.Println(s)

	m := make(map[string]bool)
	for _, v := range strings {
		AddToMapUnique(v)
	}
	fmt.Println(m)
}
