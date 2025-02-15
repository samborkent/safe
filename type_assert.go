package safe

import (
	"errors"
	"fmt"
	"reflect"
)

// TypeAssert implements safer generic type assertion. Returns zero value of the provided type if assertion fails.
// Beware that the asserted type can be nil in case of an interface, so this is not runtime safe!
func TypeAssert[T any](value any) T {
	asserted, ok := value.(T)
	if !ok {
		return *new(T)
	}

	return asserted
}

var (
	ErrPanic        = errors.New("type assertion panic")
	ErrNilInterface = errors.New("type assertion to nil interface")
	ErrTypeAssert   = errors.New("type assertion failed")
)

// TODO: test
// RequireTypeAssert implements safer type assertion with error return that can panic and recover
// in case the type cannot be asserted and the generic zero value is nil.
// Returns an error if type assertion fails. If the generic zero value is nil, panic and recover instead.
func RequireTypeAssert[T any](value any) (typ T, err error) {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if ok {
				_, _ = fmt.Printf("PANIC RECOVERED: safe.RequireTypeAssert: %s\n", err.Error())
			}
		}
	}()

	err = ErrPanic

	asserted, ok := value.(T)
	if !ok {
		err = fmt.Errorf("%w: %v cannot be asserted to %v", ErrTypeAssert, reflect.TypeOf(value), reflect.TypeOf((typ)))

		// Never return nil, panic instead
		typeOf := reflect.TypeOf(asserted)
		if typeOf == nil {
			err = fmt.Errorf("%w: %v cannot be asserted to %v", ErrNilInterface, reflect.TypeOf(value), reflect.TypeOf((typ)))
			panic(err)
		} else if typeOf.Kind() == reflect.Interface || typeOf.Kind() == reflect.Ptr {
			if reflect.ValueOf(*new(T)).IsNil() {
				panic(err)
			}
		}

		return *new(T), err
	}

	return asserted, nil
}
