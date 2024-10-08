package safe

import (
	"iter"
	"math"
	"runtime"
	"sync"
	"unsafe"
)

// A generic fixed size slice with out-of-bounds protection by clamping indices to [0, len-1].
type Array[T any] struct {
	initialized bool
	array       []T
	lock        *sync.RWMutex
}

func NewArray[T any](length int) *Array[T] {
	if length <= 0 {
		// Minimum array length is 1
		length = 1
	} else if length > math.MaxUint16 {
		// Only do a memory check for very large arrays
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		sizeOf := uint64(unsafe.Sizeof(*new(T)))

		if sizeOf == 0 {
			return &Array[T]{}
		}

		// TODO: check if this is an accurate way calculate momory available and prevent out of memory panic
		malloc := (m.Sys - m.Alloc) / sizeOf

		// Limit the length by the maximum memory available
		if malloc < math.MaxInt && length > int(malloc) {
			length = int(malloc)
		}
	}

	return &Array[T]{
		initialized: true,
		array:       make([]T, length),
		lock:        new(sync.RWMutex),
	}
}

func (a *Array[T]) Clear() {
	if !a.initialized {
		return
	}

	a.lock.Lock()
	a.array = make([]T, a.Len())
	a.lock.Unlock()
}

func (a *Array[T]) Index(index int) T {
	if !a.initialized {
		return *new(T)
	}

	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.array[a.clampIndex(index)]
}

func (a *Array[T]) Len() int {
	if !a.initialized {
		return 0
	}

	return len(a.array)
}

// TODO: test
func (a *Array[T]) Range() iter.Seq2[int, T] {
	if !a.initialized {
		return func(func(int, T) bool) {
			return
		}
	}

	return func(yield func(int, T) bool) {
		a.lock.RLock()
		defer a.lock.RUnlock()

		for i, v := range a.array {
			if !yield(i, v) {
				return
			}
		}
	}
}

func (a *Array[T]) Set(index int, value T) {
	if !a.initialized {
		return
	}

	a.lock.Lock()
	a.array[a.clampIndex(index)] = value
	a.lock.Unlock()
}

// Clamp an index to the bounds of the array
func (a *Array[T]) clampIndex(index int) int {
	if index < 0 {
		return 0
	} else if index >= a.Len() {
		return a.Len() - 1
	}

	return index
}
