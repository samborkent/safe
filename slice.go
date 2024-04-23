package safe

import (
	"runtime"
	"unsafe"
)

// TODO: test

// A generic slice with out-of-bounds protection.
// For reading, out-of-bounds indices will be clamped to [0, len-1].
// For writing, the slice will automatically grow its underlying capacity up to,
// a pre-determined maximum capacity base on the system memory statistics.
type Slice[T any] struct {
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
		slice:       make([]T, 0, capacity),
		maxCapacity: maxCapacity,
	}
}

func (s *Slice[T]) Cap() int {
	return cap(s.slice)
}

func (s *Slice[T]) Index(i int) T {
	return s.slice[s.clampIndex(i)]
}

func (s *Slice[T]) Len() int {
	return len(s.slice)
}

func (s *Slice[T]) MaxCap() int {
	return s.maxCapacity
}

func (s *Slice[T]) Set(i int, value T) {
	if i >= len(s.slice) {
		if i > s.maxCapacity {
			// If the index exceed the max capacity, grow up to the max capacity
			s.slice = append(s.slice, make([]T, s.maxCapacity-len(s.slice))...)
		} else {
			// Grow up to the index
			s.slice = append(s.slice, make([]T, i-len(s.slice)+1)...)
		}

		// Set value at last index
		s.slice[len(s.slice)-1] = value
		return
	}

	s.slice[s.clampIndex(i)] = value
}

func (s *Slice[T]) clampIndex(i int) int {
	if i < 0 {
		return 0
	} else if i >= len(s.slice) {
		return len(s.slice) - 1
	}

	return i
}
