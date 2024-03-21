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
