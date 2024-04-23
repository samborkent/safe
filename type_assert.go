package safe

import (
	"errors"
	"fmt"
	"reflect"
)

// Safer generic type assertion. Returns zero value of the provided type is assertion fails.
// Beware that the asserted type can be nil in case of an interface, so this is not runtime safe!
func TypeAssert[T comparable](value any) T {
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

// Safer type assertion with error return that can panic and recover in case the type cannot be asserted and the generic zero value is nil.
// Returns an error if type assertion fails. If the generic zero value is nil, panic and recover instead.
func RequireTypeAssert[T comparable](value any) (typ T, err error) {
	defer func() {
		if r := recover(); r != nil {
			_, _ = fmt.Printf("PANIC RECOVERED: safe.RequireTypeAssert: %s\n", r.(error).Error())
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
