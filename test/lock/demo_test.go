package demo

import (
	"math/rand"
	"testing"
	"time"
)

const TEST_TIMES = 10

func BenchmarkNormalMapLock(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(TEST_TIMES)
			NormalMapLock(k)
		}
	})

}

func BenchmarkSyncMapLock(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for p.Next() {
			k := r.Intn(TEST_TIMES)
			SyncMapLock(k)
		}
	})
}
