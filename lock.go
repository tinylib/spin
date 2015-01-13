// +build !race

// Package spin provides a
// simple spinlock.
//
package spin

// Lock spinlocks on
// an address. 0 is
// "unlocked," and 1
// is "locked."
//
//go:noescape
func Lock(l *uint32)

// Unlock unlocks a
// spinlock.
//
//go:noescape
func Unlock(l *uint32)
