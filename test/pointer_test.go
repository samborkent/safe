package safe

import (
	"testing"

	"safe"

	"github.com/stretchr/testify/assert"
)

func TestPointerDereference(t *testing.T) {
	var q *int
	_ = safe.Dereference(q)
	a := 10
	b := &a
	assert.Equal(t, a, safe.Dereference(b))
}

func TestPointerSetValue(t *testing.T) {
	a := 10
	b := &a
	assert.Equal(t, a, safe.Dereference(b))
	safe.SetValue(&b, 20)
	assert.Equal(t, 20, safe.Dereference(b))
	var c *int
	safe.SetValue(&c, 20)
	assert.Equal(t, 20, safe.Dereference(c))
	d := new(int)
	safe.SetValue(&d, 20)
	assert.Equal(t, 20, *d)
}
