package safe

import (
	"iter"
	"runtime"
	"unsafe"
)

// TODO: test

// A generic slice with out-of-bounds protection by allowing index overflow.
// For reading, out-of-bounds indices are wrapped around like in a circular buffer.
// For writing, the slice will automatically grow its underlying capacity up to,
// a pre-determined maximum capacity base on the system memory statistics.
type CircularSlice[T any] struct {
	initialized bool
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
		initialized: true,
		slice:       make([]T, 0, capacity),
		maxCapacity: maxCapacity,
	}
}

func (s CircularSlice[T]) Cap() int {
	if !s.initialized {
		return 0
	}

	return cap(s.slice)
}

func (s CircularSlice[T]) Grow(n int) {
	if !s.initialized {
		return
	}

	if n+s.Len() > s.maxCapacity {
		// If the index exceed the max capacity, grow up to the max capacity
		s.slice = append(s.slice, make([]T, s.maxCapacity-s.Len())...)
	} else {
		// Grow up to the index
		s.slice = append(s.slice, make([]T, n)...)
	}
}

func (s CircularSlice[T]) Index(index int) T {
	if !s.initialized {
		return *new(T)
	}

	return s.slice[s.wrapIndex(index)]
}

func (s CircularSlice[T]) Len() int {
	if !s.initialized {
		return 0
	}

	return len(s.slice)
}

func (s CircularSlice[T]) MaxCap() int {
	return s.maxCapacity
}

// TODO: test
func (a *CircularSlice[T]) Range() iter.Seq2[int, T] {
	if !a.initialized {
		return func(func(int, T) bool) {
			return
		}
	}

	return func(yield func(int, T) bool) {
		for i, v := range a.slice {
			if !yield(i, v) {
				return
			}
		}
	}
}

func (s CircularSlice[T]) Set(index int, value T) {
	if !s.initialized {
		return
	}

	s.slice[s.wrapIndex(index)] = value
}

func (s CircularSlice[T]) wrapIndex(index int) int {
	if index < 0 {
		return len(s.slice) + index
	} else {
		return index % len(s.slice)
	}
}
