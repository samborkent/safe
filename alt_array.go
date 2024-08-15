package safe

import (
	"iter"
	"math"
	"runtime"
	"sync/atomic"
	"unsafe"
)

// A generic fixed size slice with out-of-bounds protection by clamping indices to [0, len-1].
type AltArray[T any] struct {
	initialized bool
	array       []T
	valueLocks  []*atomic.Bool
}

func NewAltArray[T any](length int) *AltArray[T] {
	if length <= 0 {
		// Minimum array length is 1
		length = 1
	} else if length > math.MaxUint16 {
		// Only do a memory check for very large arrays
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		sizeOf := uint64(unsafe.Sizeof(*new(T)))

		if sizeOf == 0 {
			return &AltArray[T]{}
		}

		// TODO: check if this is an accurate way calculate momory available and prevent out of memory panic
		malloc := (m.Sys - m.Alloc) / sizeOf

		// Limit the length by the maximum memory available
		if malloc < math.MaxInt && length > int(malloc) {
			length = int(malloc)
		}
	}

	return &AltArray[T]{
		initialized: true,
		array:       make([]T, length),
		valueLocks:  make([]*atomic.Bool, length),
	}
}

func (a *AltArray[T]) Clear() {
	if !a.initialized {
		return
	}

	zero := *new(T)

	for i := range a.array {
		a.valueLocks[i].Store(true)
		a.array[i] = zero
		a.valueLocks[i].Store(false)
	}
}

func (a *AltArray[T]) Index(index int) T {
	if !a.initialized {
		return *new(T)
	}

	i := a.clampIndex(index)

	for {
		if !a.valueLocks[i].Load() {
			return a.array[i]
		}
	}
}

func (a *AltArray[T]) Len() int {
	if !a.initialized {
		return 0
	}

	return len(a.array)
}

// TODO: test
func (a *AltArray[T]) Range() iter.Seq2[int, T] {
	if !a.initialized {
		return func(func(int, T) bool) {
			return
		}
	}

	return func(yield func(int, T) bool) {
		for i, v := range a.array {
			a.valueLocks[i].Store(true)
			defer a.valueLocks[i].Store(false)

			if !yield(i, v) {
				return
			}
		}
	}
}

func (a *AltArray[T]) Set(index int, value T) {
	if !a.initialized {
		return
	}

	i := a.clampIndex(index)

	a.valueLocks[i].Store(true)
	a.array[i] = value
	a.valueLocks[i].Store(false)
}

// Clamp an index to the bounds of the array
func (a *AltArray[T]) clampIndex(index int) int {
	if index < 0 {
		return 0
	} else if index >= a.Len() {
		return a.Len() - 1
	}

	return index
}
