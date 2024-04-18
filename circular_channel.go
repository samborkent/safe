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
