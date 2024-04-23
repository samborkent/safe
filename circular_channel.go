package safe

import (
	"sync"
	"sync/atomic"
)

// TODO: test

// A generic non-blocking, circular channel with double closing and nil channel protection.
// Returns the zero value of the generic type if no item is available on the channel.
// Drops the oldest entry in the channel before adding a new entry to the end.
type CircularChannel[T any] struct {
	channel   chan T
	isClosed  atomic.Bool
	closeOnce sync.Once
}

func NewCircularChannel[T any](size int) *CircularChannel[T] {
	return &CircularChannel[T]{
		channel:   make(chan T, size),
		closeOnce: sync.Once{},
	}
}

func (c *CircularChannel[T]) Close() {
	c.closeOnce.Do(func() {
		c.isClosed.Store(true)
		close(c.channel)
	})
}

func (c *CircularChannel[T]) Pop() T {
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

func (c *CircularChannel[T]) Push(item T) {
	if !c.isClosed.Load() {
		select {
		case c.channel <- item:
		default:
			<-c.channel
			c.channel <- item
		}
	}
}
