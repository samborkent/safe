# safe
Toe-caps for your foot guns. [WIP]

**safe** is the friedly counterpart to **unsafe**. One of the most common issues when developing in Go are runtime panics. This package provides a set of generic functions and data structures that safeguards you from runtime panics and are meant to mitigate programmer mistakes.

## Data Structures

### Array
A generic, fixed-size array with out-of-bounds panic protection. When indexing an array *A*, the provided index *i* will be clamped to the range i E [0, len(A)-1].

### Circular Array
A generic, fixed-size circular array with out-of-bounds panic protection. When indexing a circular array, if the provided index exceeds the length of the array, it will be wrapped around. Negative index value are allowed and will also be wrapped around in the opposite direction. For example, one can index the last element of the array as: circularArray.Index(-1)

### Circular Slice
A generic, dynamically-sized array with out-of-bounds and out-of-memory panic protection.
When indexing a circular slice, if the provided index exceeds the length of the array, it will be wrapped around. Negative index value are allowed and will also be wrapped around in the opposite direction. For example, one can index the last element of the array as: circularSlice.Index(-1)
The capacity of the circular slice can be increased after initialization using the Grow method, all added values will be set to the generic zero value.
Upon initialization, an estimate is made of the maximum size in memory the slice can take based on the runtime memory profile of the system. This will be used as the maximum allowed capacity of the slice to prevent out-of-memory panics. The determined maximum safe capacity is returned by the MaxCap method.

### Channel
A generic channel with double-close and nil-channel panic protection.

### Non-Blocking Channel
A generic non-blocking channel with double-close and nil-channel panic protection. When pushing to a full channel the overflow counter is incremented. The overflow count can be retrieved using the Overflow method. When poping from an empty channel, the generic zero value is returned.

### Slice
A generic, dynamically-sized array with out-of-bounds and out-of-memory panic protection. Any index below zero will be clamped to zero. If an index is provided greater than length of the slice, the slice will automatically grow its length to that index. All values between the previous last value and the new last value will be set to the zero value of whichever type the slice is initiated with. Upon initialization, an estimate is made of the maximum size in memory the slice can take based on the runtime memory profile of the system. This will be used as the maximum allowed capacity of the slice to prevent out-of-memory panics. The determined maximum safe capacity is returned by the MaxCap method.

## Math helpers

The package also generic and safe Add, Multiply, Subtract math functions with overflow protection. There is also a safe Divide function, which provides division-by-zero panic protection. If the denominator is zero, the returned value will be the maximum value expressed by the numeric type passed to the Divide function with the sign given by the numerator.

## Pointer helpers

### Dereference
A generic function for safely dereferencing any pointer. If the pointer is nil, it will return the generic zero value.

### SetValue
Set the value that a pointer references with builtin nil-check.

## Type assertion

### TypeAssert
Safe type assertion. Returns zero value of the provided type is assertion fails. Beware that the asserted type can be nil in case of an interface.

### RequireTypeAssert
Safe type assertion with error return that can panic. Returns an error is type assertion fails. If the generic zero value is nil, panic instead.
