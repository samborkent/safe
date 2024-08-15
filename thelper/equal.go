package thelper

import (
	"reflect"
	"testing"
	"unsafe"
)

// Equal is a test helper for testing euqality of values of comparable type.
func Equal[T comparable](t *testing.T, got, want T, desc string) {
	t.Helper()

	if got != want {
		t.Errorf("%s: got %+v, want %+v", desc, got, want)
	}
}

// TODO: test
// DeepEqual tries to deduce the equality with basic reflection, and uses reflect.DeepEqual as back-up.
func DeepEqual[T any](t *testing.T, got, want T, desc string) {
	t.Helper()

	if unsafe.Sizeof(got) != unsafe.Sizeof(want) {
		t.Errorf("%s: size mismatch: got %d, want %d", desc, unsafe.Sizeof(got), unsafe.Sizeof(want))
		return
	}

	gotValue := reflect.ValueOf(got)
	wantValue := reflect.ValueOf(want)

	gotValid := gotValue.IsValid()
	wantValid := wantValue.IsValid()

	if gotValid != wantValid {
		t.Errorf("%s: validity mismatch: got %t, want %t", desc, gotValid, wantValid)
		return
	}

	gotType := gotValue.Type()
	wantType := wantValue.Type()

	if gotType != wantType {
		t.Errorf("%s: type mismatch: got %+v, want %+v", desc, gotType, wantType)
		return
	}

	gotKind := gotValue.Kind()
	wantKind := gotValue.Kind()

	if gotKind != wantKind {
		t.Errorf("%s: type kind mismatch: got %+v, want %+v", desc, gotKind, wantKind)
		return
	}

	switch gotKind {
	case reflect.Array:
		if gotValue.Len() != wantValue.Len() {
			t.Errorf("%s: array len mismatch: got %+v, want %+v", desc, gotValue.Len(), wantValue.Len())
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: array mismatch: got %+v, want %+v", desc, got, want)
		}
	case reflect.Bool:
		Equal(t, gotValue.Bool(), wantValue.Bool(), "bool mismatch")
	case reflect.Chan:
		if gotValue.Len() != wantValue.Len() {
			t.Errorf("%s: chan len mismatch: got %+v, want %+v", desc, gotValue.Len(), wantValue.Len())
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: chan mismatch: got %+v, want %+v", desc, gotValue, wantValue)
		}
	case reflect.Complex64, reflect.Complex128:
		Equal(t, gotValue.Complex(), wantValue.Complex(), "complex mismatch")
	case reflect.Float32, reflect.Float64:
		Equal(t, gotValue.Float(), wantValue.Float(), "float mismatch")
	case reflect.Func:
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: func mismatch: got %+v, want %+v", desc, gotValue, wantValue)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		Equal(t, gotValue.Int(), wantValue.Int(), "int mismatch")
	case reflect.Interface:
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: interface mismatch: got %+v, want %+v", desc, gotValue, wantValue)
		}
	case reflect.Map:
		if gotValue.IsNil() != wantValue.IsNil() {
			t.Errorf("%s: nil map mismatch: got %t, want %t", desc, gotValue.IsNil(), wantValue.IsNil())
			return
		}

		if gotValue.Len() != wantValue.Len() {
			t.Errorf("%s: map len mismatch: got %d, want %d", desc, gotValue.Len(), wantValue.Len())
			return
		}

		gotPtr := gotValue.UnsafePointer()
		wantPtr := wantValue.UnsafePointer()

		if gotPtr != wantPtr {
			t.Errorf("%s: map pointer mismatch: got %+v, want %+v", desc, gotPtr, wantPtr)
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: map mismatch: got %+v, want %+v", desc, got, want)
		}
	case reflect.Pointer:
		gotPtr := gotValue.UnsafePointer()
		wantPtr := wantValue.UnsafePointer()

		if gotPtr != wantPtr {
			t.Errorf("%s: unsafe pointer mismatch: got %+v, want %+v", desc, gotPtr, wantPtr)
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: pointer mismatch: got %+v, want %+v", desc, gotValue, wantValue)
		}
	case reflect.Slice:
		if gotValue.IsNil() != wantValue.IsNil() {
			t.Errorf("%s: nil slice mismatch: got %t, want %t", desc, gotValue.IsNil(), wantValue.IsNil())
			return
		}

		if gotValue.Len() != wantValue.Len() {
			t.Errorf("%s: map len mismatch: got %d, want %d", desc, gotValue.Len(), wantValue.Len())
			return
		}

		gotPtr := gotValue.UnsafePointer()
		wantPtr := wantValue.UnsafePointer()

		if gotPtr != wantPtr {
			t.Errorf("%s: map pointer mismatch: got %+v, want %+v", desc, gotPtr, wantPtr)
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: slice mismatch: got %+v, want %+v", desc, got, want)
		}
	case reflect.Struct:
		if gotValue.NumField() != wantValue.NumField() {
			t.Errorf("%s: struct field number mismatch: got %+v, want %+v", desc, gotValue.NumField(), wantValue.NumField())
			return
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: struct mismatch: got %+v, want %+v", desc, got, want)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		Equal(t, gotValue.Uint(), wantValue.Uint(), "uint mismatch")
	default:
		gotAny := gotValue.Interface()
		wantAny := wantValue.Interface()

		if gotAny != wantAny {
			t.Errorf("%s: uncasted interface mismatch: got %+v, want %+v", desc, gotValue, wantValue)
		}
	}
}
