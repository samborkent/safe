package safe

import (
	"sync"
	"sync/atomic"
)

type Channel[T any] struct {
	channel   chan T
	isClosed  atomic.Bool
	closeOnce sync.Once
}

// A generic channel with double closing and nil channel protection.
func NewChannel[T any](size int) *Channel[T] {
	return &Channel[T]{
		channel:   make(chan T, size),
		closeOnce: sync.Once{},
	}
}

func (c *Channel[T]) Close() {
	c.closeOnce.Do(func() {
		c.isClosed.Store(true)
		close(c.channel)
	})
}

func (c *Channel[T]) Pop() T {
	if !c.isClosed.Load() {
		return <-c.channel
	}

	return *new(T)
}

func (c *Channel[T]) Push(item T) {
	if !c.isClosed.Load() {
		c.channel <- item
	}
}
