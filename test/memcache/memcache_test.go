package memcache

import "testing"

func BenchmarkM2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m2()
	}
}

func BenchmarkM1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m1()
	}
}
