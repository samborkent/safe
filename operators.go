package safe

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Clamp[T Number](number, min, max T) T {
	if number < min {
		return min
	} else if number > max {
		return max
	}

	return number
}

func CompareFloats[T constraints.Float](a, b, tolerance, relativeTolerance T) bool {
	diff := T(math.Abs(float64(a - b)))

	if diff <= Clamp(tolerance, 0, 1) {
		return true
	} else if diff <= max(a, b)*Clamp(relativeTolerance, 0, 1) {
		return true
	}

	return false
}

func Ternary[T any](condition *bool, incase T, otherwise T) T {
	if *condition {
		return incase
	} else {
		return otherwise
	}
}
