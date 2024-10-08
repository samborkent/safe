package safe_test

import (
	"math/rand/v2"
	"testing"

	"github.com/samborkent/check"
	"github.com/samborkent/safe"
)

func TestCircularArrayIndex(t *testing.T) {
	t.Parallel()

	lenA := 4
	a := safe.NewCircularArray[int](lenA)

	a.Set(0, 1)
	a.Set(1, 2)
	a.Set(2, 3)
	a.Set(3, 4)

	check.Equal(t, a.Index(0), 1, "zero index")
	check.Equal(t, a.Index(3), 4, "last index")
	check.Equal(t, a.Index(-1), 4, "negative index")
	check.Equal(t, a.Index(4), 1, "out-of-bounds index")
}

func TestCircularArraySet(t *testing.T) {
	t.Parallel()

	x := rand.Int()
	y := rand.Int()

	lenA := 4
	a := safe.NewCircularArray[int](lenA)
	check.Equal(t, a.Index(0), 0, "zero index")

	a.Set(0, x)
	check.Equal(t, a.Index(0), x, "first set index")

	a.Set(lenA, y)
	check.Equal(t, a.Index(lenA), y, "last set index")
}
