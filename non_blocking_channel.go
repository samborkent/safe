package safe

import (
	"sync"
	"sync/atomic"
)

// TODO: test

// A generic non-blocking channel with double closing and nil channel protection.
// Returns the zero value of the generic type if no item is available on the channel.
// Increases the overflow counter when the channel is full.
type NonBlockingChannel[T any] struct {
	channel         chan T
	isClosed        atomic.Bool
	closeOnce       sync.Once
	overflowCounter uint64
	overflow        T
}

func NewNonBlockingChannel[T any](size int) *NonBlockingChannel[T] {
	return &NonBlockingChannel[T]{
		channel:   make(chan T, size),
		closeOnce: sync.Once{},
	}
}

func (c *NonBlockingChannel[T]) Close() {
	c.closeOnce.Do(func() {
		c.isClosed.Store(true)
		close(c.channel)
	})
}

func (c *NonBlockingChannel[T]) Overflow() T {
	return c.overflow
}

func (c *NonBlockingChannel[T]) OverflowCount() uint64 {
	return c.overflowCounter
}

func (c *NonBlockingChannel[T]) Pop() T {
	if !c.isClosed.Load() {
		select {
		case item := <-c.channel:
			return item
		default:
			return *new(T)
		}
	}

	return *new(T)
}

func (c *NonBlockingChannel[T]) Push(item T) {
	if !c.isClosed.Load() {
		select {
		case c.channel <- item:
		default:
			c.overflowCounter++
			c.overflow = item
		}
	}
}
