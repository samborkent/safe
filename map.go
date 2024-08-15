package safe

import (
	"iter"
	"sync"
)

// TODO: test

type Map[Key comparable, Value any] struct {
	initialized bool
	data        map[Key]Value
	lock        *sync.RWMutex
}

func NewMap[Key comparable, Value any](size int) *Map[Key, Value] {
	return &Map[Key, Value]{
		initialized: true,
		data:        make(map[Key]Value, size),
		lock:        new(sync.RWMutex),
	}
}

func (m *Map[Key, Value]) Clear(key Key) {
	if !m.initialized {
		return
	}

	m.lock.Lock()
	m.data = make(map[Key]Value, len(m.data))
	m.lock.Unlock()
}

func (m *Map[Key, Value]) Delete(key Key) {
	if !m.initialized {
		return
	}

	m.lock.Lock()
	delete(m.data, key)
	m.lock.Unlock()
}

func (m *Map[Key, Value]) Len() int {
	if !m.initialized {
		return 0
	}

	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.data)
}

func (m *Map[Key, Value]) Load(key Key) (value Value, ok bool) {
	if !m.initialized {
		return *new(Value), false
	}

	m.lock.RLock()
	value, ok = m.data[key]
	m.lock.RUnlock()
	return value, ok
}

// TODO: test
func (m *Map[Key, Value]) Range() iter.Seq2[Key, Value] {
	if !m.initialized {
		return func(func(Key, Value) bool) {
			return
		}
	}

	return func(yield func(Key, Value) bool) {
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
	previous, loaded = m.data[key]
	m.data[key] = value
	m.lock.Unlock()
	return
}

func (m *Map[Key, Value]) Store(key Key, value Value) {
	if !m.initialized {
		return
	}

	m.lock.Lock()
	m.data[key] = value
	m.lock.Unlock()
}
