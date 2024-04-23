package safe_test

import (
	"math/rand/v2"
	"testing"

	"github.com/samborkent/safe"

	"github.com/stretchr/testify/assert"
)

func TestCircularArrayIndex(t *testing.T) {
	lenA := 4
	a := safe.NewCircularArray[int](lenA)
	a.Set(0, 1)
	a.Set(1, 2)
	a.Set(2, 3)
	a.Set(3, 4)
	assert.Equal(t, 1, a.Index(0))
	assert.Equal(t, 4, a.Index(3))
	assert.Equal(t, 4, a.Index(-1))
	assert.Equal(t, 1, a.Index(4))
}

func TestCircularArraySet(t *testing.T) {
	x := rand.Int()
	y := rand.Int()

	lenA := 4
	a := safe.NewCircularArray[int](lenA)
	assert.Equal(t, 0, a.Index(0))
	a.Set(0, x)
	assert.Equal(t, x, a.Index(0))
	a.Set(lenA, y)
	assert.Equal(t, y, a.Index(0))
}
