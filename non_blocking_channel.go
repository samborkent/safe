package safe

import (
	"iter"
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

// TODO: test

// A generic non-blocking channel with double closing and nil channel protection.
// Returns the zero value of the generic type if no item is available on the channel.
// Increases the overflow counter when the channel is full.
type NonBlockingChannel[T any] struct {
	initialized     bool
	channel         chan T
	isClosed        *atomic.Bool
	closeOnce       *sync.Once
	lock            *sync.Mutex
	overflowCounter uint64
	overflow        T
}

func NewNonBlockingChannel[T any](size int) *NonBlockingChannel[T] {
	if size < 0 {
		// Minimum channel size is 0
		size = 0
	} else if size > math.MaxUint16 {
		// Only do a memory check for very large arrays
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		sizeOf := uint64(unsafe.Sizeof(*new(T)))

		if sizeOf == 0 {
			return &NonBlockingChannel[T]{}
		}

		malloc := (m.Sys - m.Alloc) / sizeOf

		// Limit the length by the maximum memory available
		if malloc < math.MaxInt && size > int(malloc) {
			size = int(malloc)
		}
	}

	return &NonBlockingChannel[T]{
		initialized: true,
		channel:     make(chan T, size),
		isClosed:    new(atomic.Bool),
		closeOnce:   new(sync.Once),
		lock:        new(sync.Mutex),
	}
}

// TODO: test
func (c *NonBlockingChannel[T]) Clear() {
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

func (c *NonBlockingChannel[T]) Close() {
	if !c.initialized {
		return
	}

	c.closeOnce.Do(func() {
		c.isClosed.Store(true)
		close(c.channel)
	})
}

func (c *NonBlockingChannel[T]) Len() int {
	if !c.initialized || c.isClosed.Load() {
		return 0
	}

	return len(c.channel)
}

func (c *NonBlockingChannel[T]) Overflow() T {
	return c.overflow
}

func (c *NonBlockingChannel[T]) OverflowCount() uint64 {
	return c.overflowCounter
}

func (c *NonBlockingChannel[T]) Pop() T {
	if !c.initialized || c.isClosed.Load() {
		return *new(T)
	}

	select {
	case item := <-c.channel:
		return item
	default:
		return *new(T)
	}
}

func (c *NonBlockingChannel[T]) Push(item T) {
	if !c.initialized || c.isClosed.Load() {
		return
	}

	select {
	case c.channel <- item:
	default:
		c.overflowCounter++
		c.overflow = item
	}
}

// TODO: test
func (c *NonBlockingChannel[T]) Range() iter.Seq[T] {
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
