package spin

import (
	"runtime"
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
