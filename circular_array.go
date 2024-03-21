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
	"math"
	"runtime"
	"unsafe"
)

// A generic fixed size slice with out-of-bounds protection by allowing index overflow.
// Out-of-bounds indices are wrapped around like in a circular buffer.
type CircularArray[T any] []T

func NewCircularArray[T any](size int) CircularArray[T] {
	if size < 0 {
		size = 0
	} else if size > math.MaxUint16 {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		sizeOf := uint64(unsafe.Sizeof(*new(T)))

		if uint64(size)*sizeOf > m.Sys-m.Alloc {
			size = int((m.Sys - m.Alloc) / sizeOf)
		}
	}

	return make(CircularArray[T], size)
}

func (a CircularArray[T]) Index(i int) T {
	return a[a.wrapIndex(i)]
}

func (a CircularArray[T]) Set(i int, value T) {
	a[a.wrapIndex(i)] = value
}

func (a CircularArray[T]) wrapIndex(i int) int {
	if i < 0 {
		return len(a) + i
	} else {
		return i % len(a)
	}
}
