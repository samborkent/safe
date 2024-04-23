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
