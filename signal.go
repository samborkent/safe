package safe

import "sync/atomic"

type Signal[T any] struct {
	initialized bool
	signal      *atomic.Pointer[T]
}

func NewSignal[T any](value T) *Signal[T] {
	ptr := &value
	signal := new(atomic.Pointer[T])
	signal.Store(ptr)

	return &Signal[T]{
		initialized: true,
		signal:      signal,
	}
}

func (s *Signal[T]) Get() T {
	if !s.initialized {
		return *new(T)
	}

	return *s.signal.Load()
}

func (s *Signal[T]) Store(value T) {
	if !s.initialized {
		return
	}

	s.signal.Store(&value)
}
