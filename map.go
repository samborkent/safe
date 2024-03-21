// Copyright 2024 Sam Borkent
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package safe

import (
	"sync"
)

// TODO: test

type Map[Key comparable, Value any] struct {
	_isInitialized bool
	data           map[Key]Value
	lock           *sync.RWMutex
}

func NewMap[Key comparable, Value any](size int) *Map[Key, Value] {
	return &Map[Key, Value]{
		_isInitialized: true,
		data:           make(map[Key]Value, size),
		lock:           new(sync.RWMutex),
	}
}

func (m *Map[Key, Value]) Delete(key Key) {
	if !m._isInitialized {
		return
	}

	m.lock.Lock()
	delete(m.data, key)
	m.lock.Unlock()
}

func (m *Map[Key, Value]) Load(key Key) (value Value, ok bool) {
	if !m._isInitialized {
		return *new(Value), false
	}

	m.lock.RLock()
	value, ok = m.data[key]
	m.lock.RUnlock()
	return value, ok
}

func (m *Map[Key, Value]) Range(f func(key Key, value Value) bool) {
	if !m._isInitialized {
		return
	}

	for k, v := range m.data {
		if !f(k, v) {
			return
		}
	}
}

func (m *Map[Key, Value]) Swap(key Key, value Value) (previous Value, loaded bool) {
	if !m._isInitialized {
		return
	}

	m.lock.Lock()
	previous, loaded = m.data[key]
	m.data[key] = value
	m.lock.Unlock()
	return
}

func (m *Map[Key, Value]) Store(key Key, value Value) {
	if !m._isInitialized {
		return
	}

	m.lock.Lock()
	m.data[key] = value
	m.lock.Unlock()
}
