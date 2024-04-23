package safe

import "reflect"

func IsNil(value any) bool {
	if value == nil {
		return true
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(value).IsNil()
	default:
		return false
	}
}
