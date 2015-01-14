// +build race !386,!amd64

package spin

import (
	"runtime"
	"sync/atomic"
)

func Lock(l *uint32) {
	for !atomic.CompareAndSwapUint32(l, 0, 1) {
		runtime.Gosched()
	}
}

func Unlock(l *uint32) {
	atomic.StoreUint32(l, 0)
}

func TryLock(l *uint32) bool {
	return atomic.CompareAndSwapUint32(l, 0, 1)
}
