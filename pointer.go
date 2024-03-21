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

// Dereference with nil pointer dereference protection.
// Returns the zero value of the underlying type is the provided pointer is nil.
func Dereference[T any](pointer *T) T {
	if pointer == nil {
		return *new(T)
	} else {
		return *pointer
	}
}

// Set the underlying value of a pointer with nil pointer dereference protection.
func SetValue[T any](pointer **T, value T) {
	if pointer != nil {
		if *pointer == nil {
			*pointer = new(T)
		}
		**pointer = value
	}
}
