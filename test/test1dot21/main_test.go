package main

import (
	"math/rand"
	"testing"
)

// 用於測試的示範數據大小
const testDataSize = 10000

// 生成隨機整數切片
func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Int()
	}
	return slice
}

// 基準測試 OriginSort 函數
func BenchmarkOriginSort(b *testing.B) {
	slice := generateRandomSlice(testDataSize)
	b.ResetTimer() // 重置計時器
	for i := 0; i < b.N; i++ {
		OriginSort(slice)
	}
}

// 基準測試 NewSort 函數
func BenchmarkNewSort(b *testing.B) {
	slice := generateRandomSlice(testDataSize)
	b.ResetTimer() // 重置計時器
	for i := 0; i < b.N; i++ {
		NewSort(slice)
	}
}

// 生成隨機 TestData 切片
func generateRandomTestDataSlice(size int) []TestData {
	slice := make([]TestData, size)
	for i := 0; i < size; i++ {
		slice[i] = TestData{Value: rand.Int()}
	}
	return slice
}

// 基準測試 OriginSort2 函數
func BenchmarkOriginSort2(b *testing.B) {
	dataSlice := generateRandomTestDataSlice(testDataSize)
	b.ResetTimer() // 重置計時器
	for i := 0; i < b.N; i++ {
		OriginSort2(dataSlice)
	}
}

// 基準測試 NewSort2 函數
func BenchmarkNewSort2(b *testing.B) {
	dataSlice := generateRandomTestDataSlice(testDataSize)
	b.ResetTimer() // 重置計時器
	for i := 0; i < b.N; i++ {
		NewSort2(dataSlice)
	}
}
