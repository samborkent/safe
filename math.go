// Copyright 2024 Sam Borkent
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package safe

import (
	"math"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// TODO: test

type Number interface {
	constraints.Integer | constraints.Float
}

// A generic add function with overflow protection.
func Add[T Number](x, y T) T {
	if y == 0 {
		return x
	}

	switch unsafe.Sizeof(T(0)) {
	case 1:
		// uint8
		if T(0)-1 > 0 {
			maxUint8 := math.MaxUint8
			if x > T(maxUint8)-y {
				return T(maxUint8)
			}
		} else {
			if y > 0 {
				// positive int8
				maxInt8 := math.MaxInt8
				if x > T(maxInt8)-y {
					return T(maxInt8)
				}
			} else {
				// negative int8
				minInt8 := math.MinInt8
				if x < T(minInt8)-y {
					return T(minInt8)
				}
			}
		}
	case 2:
		// uint16
		if T(0)-1 > 0 {
			maxUint16 := math.MaxUint16
			if x > T(maxUint16)-y {
				return T(maxUint16)
			}
		} else {
			if y > 0 {
				// positive int16
				maxInt16 := math.MaxInt16
				if x > T(maxInt16)-y {
					return T(maxInt16)
				}
			} else {
				// negative int16
				minInt16 := math.MinInt16
				if x < T(minInt16)-y {
					return T(minInt16)
				}
			}
		}
	case 4:
		// uint32
		if T(0)-1 > 0 {
			maxUint32 := math.MaxUint32
			if x > T(maxUint32)-y {
				return T(maxUint32)
			}
		} else {
			maxInt32 := math.MaxInt32
			if T(maxInt32)+1 < 0 {
				if y > 0 {
					// positive int32
					if x > T(maxInt32)-y {
						return T(maxInt32)
					}
				} else {
					// negative int32
					minInt32 := math.MinInt32
					if x < T(minInt32)-y {
						return T(minInt32)
					}
				}
			} else {
				if y > 0 {
					// positive float32
					maxFloat32 := math.MaxFloat32
					if x > T(maxFloat32)-y {
						return T(maxFloat32)
					}
				} else {
					// negative float32
					minFloat32 := -math.MaxFloat32
					if x < T(minFloat32)-y {
						return T(minFloat32)
					}
				}
			}
		}
	case 8:
		// uint64
		if T(0)-1 > 0 {
			var maxUint64 uint64 = math.MaxUint64
			if x > T(maxUint64)-y {
				return T(maxUint64)
			}
		} else {
			maxInt64 := math.MaxInt64
			if T(maxInt64)+1 < 0 {
				if y > 0 {
					// positive int64
					if x > T(maxInt64)-y {
						return T(maxInt64)
					}
				} else {
					// negative int64
					minInt64 := math.MinInt64
					if x < T(minInt64)-y {
						return T(minInt64)
					}
				}
			} else {
				if y > 0 {
					// positive float64
					maxFloat64 := math.MaxFloat64
					if x > T(maxFloat64)-y {
						return T(maxFloat64)
					}
				} else {
					// negative float64
					minFloat64 := -math.MaxFloat64
					if x < T(minFloat64)-y {
						return T(minFloat64)
					}
				}
			}
		}
	}

	return x + y
}

// A generic divide function with divide by zero protection.
// If the denominator is zero, the maximum or mininum value of the generic type will be returned based on the sign of numerator.
func Divide[T Number](x, y T) T {
	if x == 0 {
		return 0
	}

	if y == 0 {
		switch unsafe.Sizeof(T(0)) {
		case 1:
			// Check for positive overflow at zero for this type
			if T(0)-1 > 0 {
				maxUint8 := math.MaxUint8
				return T(maxUint8)
			}

			// Return min or max int value based on sign of numerator
			if x > 0 {
				maxInt8 := math.MaxInt8
				return T(maxInt8)
			} else {
				minInt8 := math.MinInt8
				return T(minInt8)
			}
		case 2:
			// Check for positive overflow at zero for this type
			if T(0)-1 > 0 {
				maxUint16 := math.MaxUint16
				return T(maxUint16)
			}

			// Return min or max int value based on sign of numerator
			if x > 0 {
				maxInt16 := math.MaxInt16
				return T(maxInt16)
			} else {
				minInt16 := math.MinInt16
				return T(minInt16)
			}
		case 4:
			// Check for positive overflow at zero for this type
			if T(0)-1 > 0 {
				maxUint32 := math.MaxUint32
				return T(maxUint32)
			}

			// Check for positive overflow at max int value
			maxInt32 := math.MaxInt32
			if T(maxInt32)+1 < 0 {
				// Return min or max int value based on sign of numerator
				if x > 0 {
					return T(maxInt32)
				} else {
					minInt32 := math.MinInt32
					return T(minInt32)
				}
			}

			// Return min or max float value based on sign of numerator
			if x > 0 {
				maxFloat32 := math.MaxFloat32
				return T(maxFloat32)
			} else {
				minFloat32 := -math.MaxFloat32
				return T(minFloat32)
			}
		case 8:
			// Check for positive overflow at zero for this type
			if T(0)-1 > 0 {
				var maxUint64 uint64 = math.MaxUint64
				return T(maxUint64)
			}

			// Check for positive overflow at max int value
			maxInt64 := math.MaxInt64
			if T(maxInt64)+1 < 0 {
				// Return min or max int value based on sign of numerator
				if x > 0 {
					return T(maxInt64)
				} else {
					minInt64 := math.MinInt64
					return T(minInt64)
				}
			}

			// Return min or max float value based on sign of numerator
			if x > 0 {
				maxFloat64 := math.MaxFloat64
				return T(maxFloat64)
			} else {
				minFloat64 := -math.MaxFloat64
				return T(minFloat64)
			}
		default:
			// Return zero for unknown type size
			return 0
		}
	}

	return x / y
}

// A generic multiply function with overflow protection.
func Multiply[T Number](x, y T) T {
	if x == 0 && y == 0 {
		return 0
	}

	switch unsafe.Sizeof(T(0)) {
	case 1:
		// uint8
		if T(0)-1 > 0 {
			maxUint8 := math.MaxUint8
			if x >= Divide(T(maxUint8), y) {
				return T(maxUint8)
			}
		} else {
			if y > 0 {
				// positive int8
				maxInt8 := math.MaxInt8
				if x >= Divide(T(maxInt8), y) {
					return T(maxInt8)
				}
			} else {
				// negative int8
				minInt8 := math.MinInt8
				if x <= Divide(T(minInt8), y) {
					return T(minInt8)
				}
			}
		}
	case 2:
		// uint16
		if T(0)-1 > 0 {
			maxUint16 := math.MaxUint16
			if x >= Divide(T(maxUint16), y) {
				return T(maxUint16)
			}
		} else {
			if y > 0 {
				// positive int16
				maxInt16 := math.MaxInt16
				if x >= Divide(T(maxInt16), y) {
					return T(maxInt16)
				}
			} else {
				// negative int16
				minInt16 := math.MinInt16
				if x <= Divide(T(minInt16), y) {
					return T(minInt16)
				}
			}
		}
	case 4:
		// uint32
		if T(0)-1 > 0 {
			maxUint32 := math.MaxUint32
			if x >= Divide(T(maxUint32), y) {
				return T(maxUint32)
			}
		} else {
			maxInt32 := math.MaxInt32
			if T(maxInt32)+1 < 0 {
				if y > 0 {
					// positive int32
					if x >= Divide(T(maxInt32), y) {
						return T(maxInt32)
					}
				} else {
					// negative int32
					minInt32 := math.MinInt32
					if x <= Divide(T(minInt32), y) {
						return T(minInt32)
					}
				}
			} else {
				if y > 0 {
					// positive float32
					maxFloat32 := math.MaxFloat32
					if x >= Divide(T(maxFloat32), y) {
						return T(maxFloat32)
					}
				} else {
					// negative float32
					minFloat32 := -math.MaxFloat32
					if x <= Divide(T(minFloat32), y) {
						return T(minFloat32)
					}
				}
			}
		}
	case 8:
		// uint64
		if T(0)-1 > 0 {
			var maxUint64 uint64 = math.MaxUint64
			if x >= Divide(T(maxUint64), y) {
				return T(maxUint64)
			}
		} else {
			maxInt64 := math.MaxInt64
			if T(maxInt64)+1 < 0 {
				if y > 0 {
					// positive int64
					if x >= Divide(T(maxInt64), y) {
						return T(maxInt64)
					}
				} else {
					// negative int64
					minInt64 := math.MinInt64
					if x <= Divide(T(minInt64), y) {
						return T(minInt64)
					}
				}
			} else {
				if y > 0 {
					// positive float64
					maxFloat64 := math.MaxFloat64
					if x >= Divide(T(maxFloat64), y) {
						return T(maxFloat64)
					}
				} else {
					// negative float64
					minFloat64 := -math.MaxFloat64
					if x <= Divide(T(minFloat64), y) {
						return T(minFloat64)
					}
				}
			}
		}
	}

	return x * y
}

// A generic subtract function with overflow protection.
func Subtract[T Number](x, y T) T {
	if y == 0 {
		return x
	}

	switch unsafe.Sizeof(T(0)) {
	case 1:
		// uint8
		if T(0)-1 > 0 {
			maxUint8 := math.MaxUint8
			if x > T(maxUint8)+y {
				return T(maxUint8)
			} else if x < y {
				return T(0)
			}
		} else {
			if y > 0 {
				// positive int8
				minInt8 := math.MinInt8
				if x < T(minInt8)+y {
					return T(minInt8)
				}
			} else {
				// negative int8
				maxInt8 := math.MaxInt8
				if x > T(maxInt8)+y {
					return T(maxInt8)
				}
			}
		}
	case 2:
		// uint16
		if T(0)-1 > 0 {
			maxUint16 := math.MaxUint16
			if x > T(maxUint16)+y {
				return T(maxUint16)
			} else if x < y {
				return T(0)
			}
		} else {
			if y > 0 {
				// positive int16
				minInt16 := math.MinInt16
				if x < T(minInt16)+y {
					return T(minInt16)
				}

			} else {
				// negative int16
				maxInt16 := math.MaxInt16
				if x > T(maxInt16)+y {
					return T(maxInt16)
				}
			}
		}
	case 4:
		// uint32
		if T(0)-1 > 0 {
			maxUint32 := math.MaxUint32
			if x > T(maxUint32)+y {
				return T(maxUint32)
			} else if x < y {
				return T(0)
			}
		} else {
			maxInt32 := math.MaxInt32
			if T(maxInt32)+1 < 0 {
				if y > 0 {
					// positive int32
					minInt32 := math.MinInt32
					if x < T(minInt32)+y {
						return T(minInt32)
					}
				} else {
					// negative int32
					if x > T(maxInt32)+y {
						return T(maxInt32)
					}
				}
			} else {
				if y > 0 {
					// positive float32
					minFloat32 := -math.MaxFloat32
					if x < T(minFloat32)+y {
						return T(minFloat32)
					}
				} else {
					// negative float32
					maxFloat32 := math.MaxFloat32
					if x > T(maxFloat32)+y {
						return T(maxFloat32)
					}
				}
			}
		}
	case 8:
		// uint64
		if T(0)-1 > 0 {
			var maxUint64 uint64 = math.MaxUint64
			if x > T(maxUint64)+y {
				return T(maxUint64)
			} else if x < y {
				return T(0)
			}
		} else {
			maxInt64 := math.MaxInt64
			if T(maxInt64)+1 < 0 {
				if y > 0 {
					// positive int64
					minInt64 := math.MinInt64
					if x < T(minInt64)+y {
						return T(minInt64)
					}
				} else {
					// negative int64
					if x > T(maxInt64)+y {
						return T(maxInt64)
					}
				}
			} else {
				if y > 0 {
					// positive float64
					minFloat64 := -math.MaxFloat64
					if x < T(minFloat64)+y {
						return T(minFloat64)
					}
				} else {
					// negative float64
					maxFloat64 := math.MaxFloat64
					if x > T(maxFloat64)+y {
						return T(maxFloat64)
					}
				}
			}
		}
	}

	return x - y
}
