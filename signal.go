package safe

type Signal[T any] struct {
	value       T
	subscribers []*T
}

func NewSignal[T any](value T) *Signal[T] {
	return &Signal[T]{
		value: value,
	}
}

func (s *Signal[T]) Get() T {
	return s.value
}

func (s *Signal[T]) Set(value T) {
	s.value = value
}

func (s *Signal[T]) Subscribe(subscriber *T) {
	s.subscribers = append(s.subscribers, subscriber)
}

func (s *Signal[T]) notify() {
}
