package spin

import (
	"runtime"
	"sync/atomic"
	"testing"
)

func init() {
	if runtime.GOMAXPROCS(0) < 2 {
		runtime.GOMAXPROCS(2)
	}
}

func TestBasicLock(t *testing.T) {
	base := uint32(0)

	Lock(&base)
	if base != 1 {
		t.Fatal("base is not 1; it is", base)
	}

	Unlock(&base)
	if base != 0 {
		t.Fatal("base is not 0; it is", base)
	}
}

func TestTryLock(t *testing.T) {
	var base uint32

	if !TryLock(&base) {
		t.Fatal("TryLock returned false on first attempt")
	}
	if base != 1 {
		t.Fatal("base should be 1")
	}
	if TryLock(&base) {
		t.Fatal("TryLock succeeded on second attempt")
	}
}

func TestConcurrentLock(t *testing.T) {
	base := uint32(0)
	Lock(&base)
	go Unlock(&base) // remote unlock
	Lock(&base)
}

func BenchmarkUncontested(b *testing.B) {
	var l uint32
	for i := 0; i < b.N; i++ {
		Lock(&l)
		Unlock(&l)
	}
}

func BenchmarkSyncAtomicUncontested(b *testing.B) {
	var l uint32
	for i := 0; i < b.N; i++ {
		for !atomic.CompareAndSwapUint32(&l, 0, 1) {
		}
		atomic.StoreUint32(&l, 0)
	}
}
