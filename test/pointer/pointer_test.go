package pointer

import (
	"fmt"
	"testing"
)

func ExampleGetAddrReflect() {
	a := map[string]any{
		"A": "A",
	}
	fmt.Println(GetAddrReflect(&a))
	fmt.Println(GetAddrReflect2(&a))

	// output:
	// a
	// b
}

func BenchmarkGetAddrReflect(b *testing.B) {
	//x := generateIntSlice(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetAddrReflect(1)
	}
}

func BenchmarkGetAddrReflect2(b *testing.B) {
	//x := generateIntSlice(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetAddrReflect2(1)
	}
}

func BenchmarkGetAddrFmt(b *testing.B) {
	//x := generateIntSlice(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetAddrFmt(1)
	}
}
