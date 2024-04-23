package safe

import (
	"math"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

// TODO: test

// A generic, circular channel that blocks on pop, but not on push,
//
//	with double closing and nil channel protection.
//
// Block until an new entry is available on pop.
// If the channel is full, the oldest entry will be dropped on push.
type CircularChannel[T any] struct {
	initialized bool
	channel     chan T
	isClosed    *atomic.Bool
	closeOnce   *sync.Once
	lock        *sync.Mutex
}

func NewCircularChannel[T any](size int) *CircularChannel[T] {
	if size < 0 {
		// Minimum channel size is 0
		size = 0
	} else if size > math.MaxUint16 {
		// Only do a memory check for very large arrays
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		sizeOf := uint64(unsafe.Sizeof(*new(T)))

		if sizeOf == 0 {
			return &CircularChannel[T]{}
		}

		malloc := (m.Sys - m.Alloc) / sizeOf

		// Limit the length by the maximum memory available
		if malloc < math.MaxInt && size > int(malloc) {
			size = int(malloc)
		}
	}

	return &CircularChannel[T]{
		initialized: true,
		channel:     make(chan T, size),
		isClosed:    new(atomic.Bool),
		closeOnce:   new(sync.Once),
		lock:        new(sync.Mutex),
	}
}

// TODO: test
func (c *CircularChannel[T]) Clear() {
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

func (c *CircularChannel[T]) Close() {
	if !c.initialized {
		return
	}

	c.closeOnce.Do(func() {
		c.isClosed.Store(true)
		close(c.channel)
	})
}

func (c *CircularChannel[T]) Len() int {
	if !c.initialized || c.isClosed.Load() {
		return 0
	}

	return len(c.channel)
}

func (c *CircularChannel[T]) Pop() T {
	if !c.initialized || c.isClosed.Load() {
		return *new(T)
	}

	return <-c.channel
}

func (c *CircularChannel[T]) Push(item T) {
	if !c.initialized || c.isClosed.Load() {
		return
	}

	select {
	case c.channel <- item:
	default:
		<-c.channel
		c.channel <- item
	}
}
