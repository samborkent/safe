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
	"runtime"
	"unsafe"
)

// TODO: test

// A generic slice with out-of-bounds protection by allowing index overflow.
// For reading, out-of-bounds indices are wrapped around like in a circular buffer.
// For writing, the slice will automatically grow its underlying capacity up to,
// a pre-determined maximum capacity base on the system memory statistics.
type CircularSlice[T any] struct {
	slice       []T
	maxCapacity int
}

func NewCircularSlice[T any](capacity int) CircularSlice[T] {
	if capacity < 0 {
		capacity = 0
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	maxCapacity := int((m.Sys - m.Alloc) / uint64(unsafe.Sizeof(*new(T))))

	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	return CircularSlice[T]{
		slice:       make([]T, 0, capacity),
		maxCapacity: maxCapacity,
	}
}

func (s CircularSlice[T]) Cap() int {
	return cap(s.slice)
}

func (s CircularSlice[T]) Grow(n int) {
	if n+s.Len() > s.maxCapacity {
		// If the index exceed the max capacity, grow up to the max capacity
		s.slice = append(s.slice, make([]T, s.maxCapacity-s.Len())...)
	} else {
		// Grow up to the index
		s.slice = append(s.slice, make([]T, n)...)
	}
}

func (s CircularSlice[T]) Index(i int) T {
	return s.slice[s.wrapIndex(i)]
}

func (s CircularSlice[T]) Len() int {
	return len(s.slice)
}

func (s CircularSlice[T]) MaxCap() int {
	return s.maxCapacity
}

func (s CircularSlice[T]) Set(i int, value T) {
	s.slice[s.wrapIndex(i)] = value
}

func (s CircularSlice[T]) wrapIndex(i int) int {
	if i < 0 {
		return len(s.slice) + i
	} else {
		return i % len(s.slice)
	}
}
