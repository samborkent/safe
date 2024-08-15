package safe

import (
	"iter"
	"runtime"
	"unsafe"
)

// TODO: test

// A generic slice with out-of-bounds protection.
// For reading, out-of-bounds indices will be clamped to [0, len-1].
// For writing, the slice will automatically grow its underlying capacity up to,
// a pre-determined maximum capacity base on the system memory statistics.
type Slice[T any] struct {
	initialized bool
	slice       []T
	maxCapacity int
}

func NewSlice[T any](capacity int) *Slice[T] {
	if capacity < 0 {
		capacity = 0
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	maxCapacity := int((m.Sys - m.Alloc) / uint64(unsafe.Sizeof(*new(T))))

	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	return &Slice[T]{
		initialized: true,
		slice:       make([]T, 0, capacity),
		maxCapacity: maxCapacity,
	}
}

func (s *Slice[T]) Cap() int {
	if !s.initialized {
		return 0
	}

	return cap(s.slice)
}

func (s *Slice[T]) Index(index int) T {
	if !s.initialized {
		return *new(T)
	}

	return s.slice[s.clampIndex(index)]
}

func (s *Slice[T]) Len() int {
	if !s.initialized {
		return 0
	}

	return len(s.slice)
}

func (s *Slice[T]) MaxCap() int {
	return s.maxCapacity
}

// TODO: test
func (a *Slice[T]) Range() iter.Seq2[int, T] {
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

func (s *Slice[T]) Set(index int, value T) {
	if !s.initialized {
		return
	}

	if index >= len(s.slice) {
		if index > s.maxCapacity {
			// If the index exceed the max capacity, grow up to the max capacity
			s.slice = append(s.slice, make([]T, s.maxCapacity-len(s.slice))...)
		} else {
			// Grow up to the index
			s.slice = append(s.slice, make([]T, index-len(s.slice)+1)...)
		}

		// Set value at last index
		s.slice[len(s.slice)-1] = value
		return
	}

	s.slice[s.clampIndex(index)] = value
}

func (s *Slice[T]) clampIndex(index int) int {
	if index < 0 {
		return 0
	} else if index >= len(s.slice) {
		return len(s.slice) - 1
	}

	return index
}
