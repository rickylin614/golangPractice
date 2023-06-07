package demo

import (
	"sync"
	"time"
)

var (
	memberLock   = map[int]struct{}{}
	memberLocker = &sync.Mutex{}
	memberSync   = &sync.Map{}
)

func NormalMapLock(s int) {
	memberLocker.Lock()
	if _, ok := memberLock[s]; ok {
		memberLocker.Unlock()
		time.Sleep(time.Nanosecond * 10)
		return
	} else {
		memberLock[s] = struct{}{}
	}
	memberLocker.Unlock()
	defer func() {
		memberLocker.Lock()
		memberLocker.Unlock()
	}()
	time.Sleep(time.Nanosecond * 10)
}

func SyncMapLock(s int) {
	if _, ok := memberSync.LoadOrStore(s, struct{}{}); ok {
		time.Sleep(time.Nanosecond * 10)
		return
	}
	defer memberSync.Delete(s)
	time.Sleep(time.Nanosecond * 10)
}
