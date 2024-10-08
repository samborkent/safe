package safe_test

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/samborkent/check"
	"github.com/samborkent/safe"
)

func TestNewArray(t *testing.T) {
	t.Parallel()

	lenA := 6
	a := safe.NewArray[float64](lenA)
	check.Equal(t, a.Len(), lenA, "random length")

	b := safe.NewArray[float64](0)
	check.Equal(t, b.Len(), 1, "zero length")

	c := safe.NewArray[float64](-1)
	check.Equal(t, c.Len(), 1, "negative length")

	lenD := math.MaxInt64
	d := safe.NewArray[float64](lenD)
	check.Less(t, d.Len(), lenD, "max length")
}

func TestArrayIndex(t *testing.T) {
	t.Parallel()

	lenA := 4
	a := safe.NewArray[int](lenA)

	a.Set(0, 1)
	a.Set(1, 2)
	a.Set(2, 3)
	a.Set(3, 4)

	check.Equal(t, a.Index(0), 1, "zero index")
	check.Equal(t, a.Index(lenA-1), 4, "last index")
	check.Equal(t, a.Index(-1), 1, "negative index")
	check.Equal(t, a.Index(lenA), 4, "out-of-bounds index")
}

func TestArraySet(t *testing.T) {
	t.Parallel()

	x := rand.Int()
	y := rand.Int()

	lenA := 4
	a := safe.NewArray[int](lenA)
	check.Equal(t, a.Index(0), 0, "unset zero element")

	a.Set(0, x)
	check.Equal(t, a.Index(0), x, "set zero element")

	a.Set(lenA, y)
	check.Equal(t, a.Index(lenA-1), y, "set last element")
}
