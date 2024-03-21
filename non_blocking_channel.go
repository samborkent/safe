// Copyright 2024 Sam Borkent
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func (c *NonBlockingChannel[T]) Overflow() uint64 {
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
		}
	}
}
