package thelper

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func Less[T constraints.Ordered](t *testing.T, got, want T, desc string) {
	t.Helper()

	if got >= want {
		t.Errorf("%s: %v should be less then %v", desc, got, want)
	}
}
