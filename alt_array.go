package safe

import (
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

func (a *AltArray[T]) Index(i int) T {
	if !a.initialized {
		return *new(T)
	}

	index := a.clampIndex(i)

	for {
		if !a.valueLocks[index].Load() {
			return a.array[index]
		}
	}
}

func (a *AltArray[T]) Len() int {
	if !a.initialized {
		return 0
	}

	return len(a.array)
}

func (a *AltArray[T]) Set(i int, value T) {
	if !a.initialized {
		return
	}

	index := a.clampIndex(i)

	a.valueLocks[index].Store(true)
	a.array[index] = value
	a.valueLocks[index].Store(false)
}

// Clamp an index to the bounds of the array
func (a *AltArray[T]) clampIndex(i int) int {
	if i < 0 {
		return 0
	} else if i >= a.Len() {
		return a.Len() - 1
	}

	return i
}
