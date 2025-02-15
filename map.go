package safe

import (
	"iter"
	"sync"
)

// TODO: test

type Map[Key comparable, Value any] struct {
	lock        sync.RWMutex
	data        map[Key]Value
	initialized bool
}

func NewMap[Key comparable, Value any](size int) *Map[Key, Value] {
	return &Map[Key, Value]{
		initialized: true,
		data:        make(map[Key]Value, size),
	}
}

func (m *Map[Key, Value]) Clear(key Key) {
	if !m.initialized {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	m.data = make(map[Key]Value, len(m.data))
}

func (m *Map[Key, Value]) Delete(key Key) {
	if !m.initialized {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.data, key)
}

func (m *Map[Key, Value]) Len() int {
	if !m.initialized {
		return 0
	}

	return len(m.data)
}

func (m *Map[Key, Value]) Load(key Key) (value Value, ok bool) {
	if !m.initialized {
		return *new(Value), false
	}

	m.lock.RLock()
	defer m.lock.RUnlock()

	value, ok = m.data[key]
	return value, ok
}

// TODO: test
func (m *Map[Key, Value]) Range() iter.Seq2[Key, Value] {
	if !m.initialized {
		return func(func(Key, Value) bool) {}
	}

	return func(yield func(Key, Value) bool) {
		m.lock.RLock()
		defer m.lock.RUnlock()

		for k, v := range m.data {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (m *Map[Key, Value]) Swap(key Key, value Value) (previous Value, loaded bool) {
	if !m.initialized {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	previous, loaded = m.data[key]
	m.data[key] = value

	return previous, loaded
}

func (m *Map[Key, Value]) Store(key Key, value Value) {
	if !m.initialized {
		return
	}

	m.lock.Lock()
	defer m.lock.Unlock()

	m.data[key] = value
}
