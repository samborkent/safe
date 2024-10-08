package safe

import (
	"math/rand/v2"
	"testing"

	"github.com/samborkent/check"
	"github.com/samborkent/safe"
)

func TestPointerDereference(t *testing.T) {
	var q *int
	_ = safe.Dereference(q)
	a := 10
	b := &a
	check.Equal(t, safe.Dereference(b), a, "dereference pointer")
}

func TestPointerSetValue(t *testing.T) {
	a := rand.Int()
	b := &a
	check.Equal(t, safe.Dereference(b), a, "dereference pointer")

	val1 := rand.Int()
	safe.SetValue(&b, val1)
	check.Equal(t, safe.Dereference(b), val1, "dereference set pointer")

	val2 := rand.Int()
	var c *int
	safe.SetValue(&c, val2)
	check.Equal(t, safe.Dereference(c), val2, "dereference new set pointer")

	val3 := rand.Int()
	d := new(int)
	safe.SetValue(&d, val3)
	check.Equal(t, *d, val3, "unsafely dereference new set pointer")
}
