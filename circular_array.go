package safe

import (
	"math"
	"runtime"
	"sync"
	"unsafe"
)

// A generic fixed size slice with out-of-bounds protection by allowing index overflow.
// Out-of-bounds indices are wrapped around like in a circular buffer.
type CircularArray[T any] struct {
	initialized bool
	array       []T
	lock        *sync.RWMutex
}

func NewCircularArray[T any](length int) *CircularArray[T] {
	if length <= 0 {
		// Minimum array length is 1
		length = 1
	} else if length > math.MaxUint16 {
		// Only do a memory check for very large arrays
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		sizeOf := uint64(unsafe.Sizeof(*new(T)))

		if sizeOf == 0 {
			return &CircularArray[T]{}
		}

		// TODO: check if this is an accurate way calculate momory available and prevent out of memory panic
		malloc := (m.Sys - m.Alloc) / sizeOf

		// Limit the length by the maximum memory available
		if malloc < math.MaxInt && length > int(malloc) {
			length = int(malloc)
		}
	}

	return &CircularArray[T]{
		initialized: true,
		array:       make([]T, length),
		lock:        new(sync.RWMutex),
	}
}

func (a *CircularArray[T]) Clear() {
	if !a.initialized {
		return
	}

	a.lock.Lock()
	a.array = make([]T, a.Len())
	a.lock.Unlock()
}

func (a *CircularArray[T]) Index(i int) T {
	if !a.initialized {
		return *new(T)
	}

	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.array[a.wrapIndex(i)]
}

func (a *CircularArray[T]) Len() int {
	if !a.initialized {
		return 0
	}

	a.lock.RLock()
	defer a.lock.RUnlock()
	return len(a.array)
}

func (a *CircularArray[T]) Set(i int, value T) {
	if !a.initialized {
		return
	}

	a.lock.Lock()
	a.array[a.wrapIndex(i)] = value
	a.lock.Unlock()
}

func (a *CircularArray[T]) wrapIndex(i int) int {
	if i < 0 {
		return a.Len() + i
	} else {
		return i % a.Len()
	}
}
