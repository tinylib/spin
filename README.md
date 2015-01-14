spin
=====

Spin provides a simple spinlock.

### Usage

Since goroutines blocking on a spinlock don't accomplish any useful work while they are blocked (as opposed to goroutines blocked on a `sync.Mutex`, which yield to runnable goroutines), a spinlock should only be used to protect very small operations.

Here's an example of a goroutine-safe data structure that keeps track of two integers that are always written and read atomically:

```go
type Container struct {
	one  int
	two  int
	lock uint32
}

func (c *Container) Get() (one int, two int) {
	spin.Lock(&c.lock)
	one, two = c.one, c.two
	spin.Unlock(&c.lock)
	return
}

func (c *Container) Set(one int, two int) {
	spin.Lock(&c.lock)
	c.one, c.two = one, two
	spin.Unlock(&c.lock)
}
```

Since we're using the lock to protect only two loads/stores, the overhead of yielding a goroutine to the scheduler is large compared to the time of the operation, and thus a spinlock is a more efficient way to serialize access to the container's fields.