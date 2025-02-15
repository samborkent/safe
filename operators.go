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

func ClampMin[T Number](number, min T) T {
	if number < min {
		return min
	}

	return number
}

func ClampMax[T Number](number, max T) T {
	if number > max {
		return max
	}

	return number
}

// CompareFloats compares two floats. First, it tries to compare the floats exactly.
// Secondly, it compared them using the absolute tolerance if tolerance >= 0.
// Lastly, it compares the floats using the relative tolerance if relativeTolerance >= 0.
func CompareFloats[T constraints.Float](a, b, tolerance, relativeTolerance T) bool {
	if a == b {
		return true
	}

	if tolerance < 0 || relativeTolerance < 0 {
		return false
	}

	diff := T(math.Abs(float64(a - b)))

	if tolerance != 0 && diff <= Clamp(tolerance, 0, 1) {
		return true
	}

	if relativeTolerance != 0 && diff <= max(a, b)*Clamp(relativeTolerance, 0, 1) {
		return true
	}

	return false
}

// Ternary functions as a ternary ? operator found in functional programming.
// If the condition is true, it returns the first case, otherwise the secone case.
func Ternary[T any](condition bool, incase, otherwise T) T {
	if condition {
		return incase
	} else {
		return otherwise
	}
}
