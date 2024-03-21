package safe_test

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/samborkent/safe"

	"github.com/stretchr/testify/assert"
)

func TestNewArray(t *testing.T) {
	lenA := 6
	a := safe.NewArray[float64](lenA)
	assert.Equal(t, lenA, len(a))

	lenB := 0
	b := safe.NewArray[float64](lenB)
	assert.Equal(t, lenB, len(b))

	lenC := -1
	c := safe.NewArray[float64](lenC)
	assert.Equal(t, 0, len(c))

	lenD := math.MaxInt64
	_ = safe.NewArray[float64](lenD)
}

func TestArrayIndex(t *testing.T) {
	lenA := 4
	a := safe.NewArray[int](lenA)
	_ = a.Index(0)
	_ = a.Index(lenA)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	assert.Equal(t, 1, a.Index(-1))
	assert.Equal(t, 4, a.Index(4))
}

func TestArraySet(t *testing.T) {
	x := rand.Int()
	y := rand.Int()

	lenA := 4
	a := safe.NewArray[int](lenA)
	assert.Equal(t, 0, a.Index(0))
	a.Set(0, x)
	assert.Equal(t, x, a.Index(0))
	a.Set(lenA, y)
	assert.Equal(t, y, a.Index(lenA-1))
}
