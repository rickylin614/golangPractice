package slogtest

import "testing"

func BenchmarkSlogPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlogPrint()
	}
}

func BenchmarkZlogPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZlogPrint()
	}
}

func BenchmarkZlogSugarPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZlogSugarPrint()
	}
}
