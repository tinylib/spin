package spin

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestBasicLock(t *testing.T) {
	var base uint32

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
	var base uint32
	Lock(&base)
	go func() {
		time.Sleep(1 * time.Microsecond)
		Unlock(&base) // remote unlock
	}()
	Lock(&base)
	Unlock(&base)

	incr := 0
	ninc := 100000
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for j := 0; j < ninc; j++ {
			Lock(&base)
			incr++
			Unlock(&base)
		}
		wg.Done()
	}()
	for i := 0; i < ninc; i++ {
		Lock(&base)
		incr++
		Unlock(&base)
	}
	wg.Wait()
	want := 2 * ninc
	if incr != want {
		t.Fatal("want", want, "but got", incr)
	}
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

func BenchmarkParallelAdd(b *testing.B) {
	var l uint32
	var i uint64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Lock(&l)
			i++
			Unlock(&l)
		}
	})
}

func BenchmarkAtomicParallelAdd(b *testing.B) {
	var l uint32
	var i uint64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for !atomic.CompareAndSwapUint32(&l, 0, 1) {
			}
			i++
			atomic.StoreUint32(&l, 0)
		}
	})
}

func BenchmarkMutexParallelAdd(b *testing.B) {
	lock := new(sync.Mutex)
	var i uint64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			i++
			lock.Unlock()
		}
	})
}
