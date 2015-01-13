// +build !race

// Package spin provides a
// simple spinlock.
//
// Since spinlocks continue
// to use CPU resources while
// they attempt to acquire a lock,
// they should only be used to protect
// extremely small (read: fast) regions of code,
// particularly in the case where that code is
// usually uncontested.
//
// It is also possible to use
// TryLock() to do useful work while
// waiting to acquire a lock. In pseudo-code:
//
// 	for !spin.TryLock(&lock) {
//		doOtherWork()
//  }
//  doCriticalWork()
//  spin.Unlock(&lock)
//
package spin

// Lock spinlocks on
// an address. 0 is
// "unlocked," and 1
// is "locked."
//
//go:noescape
func Lock(l *uint32)

// TryLock attempts
// to atomically
// change the value
// of *l from 0 to 1,
// and returns whether or
// not it was successful.
//
//go:noescape
func TryLock(l *uint32) bool

// Unlock unlocks a
// spinlock.
//
//go:noescape
func Unlock(l *uint32)
