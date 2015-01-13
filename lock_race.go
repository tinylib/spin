// +build race !386,!amd64

package spin

import "sync/atomic"

func Lock(l *uint32) {
	for !atomic.CompareAndSwapUint32(l, 0, 1) {
	}
}

func Unlock(l *uint32) {
	atomic.StoreUint32(l, 0)
}
