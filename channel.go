package safe

import (
	"iter"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

// A generic blocking channel with double closing and nil channel protection.
type Channel[T any] struct {
	initialized bool
	channel     chan T
	isClosed    *atomic.Bool
	closeOnce   *sync.Once
	lock        *sync.Mutex
}

// A generic channel with double closing and nil channel protection.
func NewChannel[T any](size int) *Channel[T] {
	if size < 0 {
		// Minimum channel size is 0
		size = 0
	} else if size > math.MaxUint16 {
		// Only do a memory check for very large arrays
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		sizeOf := uint64(unsafe.Sizeof(*new(T)))

		if sizeOf == 0 {
			return &Channel[T]{}
		}

		malloc := (m.Sys - m.Alloc) / sizeOf

		// Limit the length by the maximum memory available
		if malloc < math.MaxInt && size > int(malloc) {
			size = int(malloc)
		}
	}

	return &Channel[T]{
		initialized: true,
		channel:     make(chan T, size),
		isClosed:    new(atomic.Bool),
		closeOnce:   new(sync.Once),
		lock:        new(sync.Mutex),
	}
}

// TODO: test
func (c *Channel[T]) Clear() {
	if !c.initialized {
		return
	}

	c.lock.Lock()

	if !c.isClosed.Load() {
		close(c.channel)
	}

	c.channel = make(chan T, len(c.channel))
	c.lock.Unlock()
}

func (c *Channel[T]) Close() {
	if !c.initialized {
		return
	}

	// TODO: check if this is necessary and if the atomic bool is not sufficient
	c.closeOnce.Do(func() {
		c.isClosed.Store(true)
		close(c.channel)
	})
}

func (c *Channel[T]) Len() int {
	if !c.initialized || c.isClosed.Load() {
		return 0
	}

	return len(c.channel)
}

func (c *Channel[T]) Pop() T {
	if !c.initialized || c.isClosed.Load() {
		return *new(T)
	}

	return <-c.channel
}

func (c *Channel[T]) Push(item T) {
	if !c.initialized || c.isClosed.Load() {
		return
	}

	c.channel <- item
}

// TODO: test
func (c *Channel[T]) Range() iter.Seq[T] {
	if !c.initialized || c.isClosed.Load() {
		return func(func(T) bool) {
			return
		}
	}

	return func(yield func(T) bool) {
		for v := range c.channel {
			if !yield(v) {
				return
			}
		}
	}
}
