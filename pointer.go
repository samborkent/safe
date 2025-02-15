package safe

// Dereference can be used to safely dereference a pointer with nil pointer dereference protection.
// Returns the zero value of the underlying type if the provided pointer is nil.
func Dereference[T any](pointer *T) T {
	if pointer == nil {
		return *new(T)
	} else {
		return *pointer
	}
}

// SetValue sets the underlying value of a pointer with nil pointer dereference protection.
// No-op if the underlying pointer is nil.
func SetValue[T any](pointer **T, value T) {
	if pointer != nil {
		if *pointer == nil {
			*pointer = new(T)
		}
		**pointer = value
	}
}
