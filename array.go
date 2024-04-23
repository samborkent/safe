package safe

import (
	"math"
	"runtime"
	"unsafe"
)

// A generic fixed size slice with out-of-bounds protection by clamping indices to [0, len-1].
type Array[T any] struct {
	initialized bool
	array       []T
}

func NewArray[T any](length int) *Array[T] {
	if length < 0 {
		length = 0
	} else if length > math.MaxUint16 {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		sizeOf := uint64(unsafe.Sizeof(*new(T)))

		if uint64(length)*sizeOf > m.Sys-m.Alloc {
			length = int((m.Sys - m.Alloc) / sizeOf)
		}
	}

	return &Array[T]{
		initialized: true,
		array:       make([]T, length),
	}
}

func (a Array[T]) Index(i int) T {
	if !a.initialized {
		return *new(T)
	}

	return a.array[a.clampIndex(i)]
}

func (a Array[T]) Len() int {
	if !a.initialized {
		return 0
	}

	return len(a.array)
}

func (a Array[T]) Set(i int, value T) {
	if !a.initialized {
		return
	}

	a.array[a.clampIndex(i)] = value
}

func (a Array[T]) clampIndex(i int) int {
	if i < 0 {
		return 0
	} else if i >= a.Len() {
		return a.Len() - 1
	}

	return i
}
