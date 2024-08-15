# safe

Toe-caps for your foot guns. [WIP]

`safe` is the friedly counterpart to `unsafe`. One of the most common issues when developing in Go are runtime panics. This package provides a set of generic functions and data structures that safeguards you from runtime panics and are meant to mitigate programmer mistakes. Now updated to Go 1.23 to support iterators.

## Data Structures

All iterable data structures provide a `Range` iterator method which can be used with the `range` expression, introduced in Go 1.23.

### Array

`Array` is a generic, fixed-size array with out-of-bounds panic protection. When indexing an array `A`, the provided index `i` will be clamped to the range `i E [0, len(A)-1]`.

### CircularArray

`CircularArray` is a generic, fixed-size circular array with out-of-bounds panic protection.

When indexing a circular array, if the provided index exceeds the length of the array, it will be wrapped around. Negative index value are allowed and will also be wrapped around in the opposite direction. For example, one can index the last element of the array as:

```
lastElement := circularArray.Index(-1)
```

### CircularSlice

`CircularSlice` is a generic, dynamically-sized array with out-of-bounds and out-of-memory panic protection.

When indexing a circular slice, if the provided index exceeds the length of the array, it will be wrapped around. Negative index value are allowed and will also be wrapped around in the opposite direction. See [CircularArray](#circular-array) for an example.

The capacity of the circular slice can be increased after initialization using the Grow method, all added values will be set to the generic zero value.
Upon initialization, an estimate is made of the maximum size in memory the slice can take based on the runtime memory profile of the system. This will be used as the maximum allowed capacity of the slice to prevent out-of-memory panics. The determined maximum safe capacity is returned by the MaxCap method.

### Channel

`Channel` is a generic channel with double-close and nil-channel panic protection.

### NonBlockingChannel

`NonBlockingChannel` is a generic non-blocking channel with double-close and nil-channel panic protection.

When pushing to a full channel the overflow counter is incremented. The overflow count can be retrieved using the Overflow method. When poping from an empty channel, the generic zero value is returned.

### Map

`Map` is a generic map with nil-panic protection and a built-in lock. The key must be a `comparable` type.

### Slice

`Slice` is a generic, dynamically-sized array with out-of-bounds and out-of-memory panic protection.

Any index below zero will be clamped to zero. If an index is provided greater than length of the slice, the slice will automatically grow its length to that index. All values between the previous last value and the new last value will be set to the zero value of whichever type the slice is initiated with. Upon initialization, an estimate is made of the maximum size in memory the slice can take based on the runtime memory profile of the system. This will be used as the maximum allowed capacity of the slice to prevent out-of-memory panics. The determined maximum safe capacity is returned by the MaxCap method.

## Math helpers

### Add, Multiply, Subtract

`Add`, `Multiply`, and `Subtract` are generic math functions that offer overflow protection.

### Clamp, ClampMin, ClampMax

`Clamp` is a generic function to clamp any number between a minimum and maximum value. `ClampMin` and `ClampMax` work the same but only clamp the lower or upper limit.

### CompareFloats

`CompareFloats` is a safe way to compare two floating point numbers providing an absolute and a relative tolerance.

### Divide

`Divide` is a generic division function that provides division-by-zero panic protection. When the denominator is zero, the returned value will be the maximum value expressed by the numeric type passed to the Divide function with the sign given by the numerator.

### Ternary

`Ternary` can be used as a ternary operator for inline comparisons.

## Pointer helpers

### Dereference

`Dereference` is a generic function for safely dereferencing any pointer. If the pointer is nil, it will return the generic zero value.

### SetValue

`SetValue` is a generic function that sets the value of a pointer reference with a builtin nil-check.

## Type assertion

### TypeAssert

`TypeAssert` is a safe type assertion function. It returns the zero value of the provided type if the assertion fails. Beware that the asserted type can be nil in case of an interface.

### RequireTypeAssert

`RequireTypeAssert` is a type assertion assetion function with error return that can panic. Returns an error is type assertion fails. Instead of returning a nil-value when it asserts an interface as nil, like [TypeAssert](#type-assert) does, it will panic and recover instead.

## Miscellaneous

### IsNil

`IsNil` is a helper function to do a deep `nil` check on `any` type.

## Questions

### Why do the math functions use the unsafe package?

At the moment (Go 1.23), Go does not offer generic type assertion or generic type switching. So, after extensive testing and becnhmarking, I found that switching of the byte size of the generic zero type is currently the fastest way to differentiate between different numeric types. However, this does require som examination of the types overflow behaviour to deduce the concrete type See the implementations of `Add`, `Divide`, `Multiply`, and `Subtract` as example.

### Why do you dereference `new(T)` without nil-checking?

Currently (Go 1.23), there is no pre-defined way to return the zero value of a generic type, but you can instantiate a pointer to a generic zero type using `new(T)`. This value is always safe to dereference and will not panic, so we can use `*new(T)`. However, this can return `nil` in some cases, so package users are still expected to do proper nil-checking when working with nullable types.
