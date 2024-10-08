package safe_test

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/samborkent/check"
	"github.com/samborkent/safe"
)

func TestSubtract(t *testing.T) {
	u8 := uint8(math.MaxUint8)
	check.Equal(t, safe.Subtract(0, u8), 0, "uint8")

	u16 := uint16(math.MaxUint16)
	check.Equal(t, safe.Subtract(0, u16), 0, "uint16")

	u32 := uint32(math.MaxUint32)
	check.Equal(t, safe.Subtract(0, u32), 0, "uint32")

	u64 := uint64(math.MaxUint64)
	check.Equal(t, safe.Subtract(0, u64), 0, "uint64")

	u := uint(math.MaxUint)
	check.Equal(t, safe.Subtract(0, u), 0, "uint")

	p8, n8 := int8(math.MaxInt8), int8(math.MinInt8)
	check.Equal(t, safe.Subtract(p8, n8), p8, "int8 positive")
	check.Equal(t, safe.Subtract(n8, p8), n8, "int8 negative")

	p16, n16 := int16(math.MaxInt16), int16(math.MinInt16)
	check.Equal(t, safe.Subtract(p16, n16), p16, "int16 positive")
	check.Equal(t, safe.Subtract(n16, p16), n16, "int16 negative")

	p32, n32 := int32(math.MaxInt32), int32(math.MinInt32)
	check.Equal(t, safe.Subtract(p32, n32), p32, "int32 positive")
	check.Equal(t, safe.Subtract(n32, p32), n32, "int32 negative")

	p64, n64 := int64(math.MaxInt64), int64(math.MinInt64)
	check.Equal(t, safe.Subtract(p64, n64), p64, "int64 positive")
	check.Equal(t, safe.Subtract(n64, p64), n64, "int64 negative")

	pi, ni := int(math.MaxInt), int(math.MinInt)
	check.Equal(t, safe.Subtract(pi, ni), pi, "int positive")
	check.Equal(t, safe.Subtract(ni, pi), ni, "int negative")

	pf32, nf32 := float32(math.MaxFloat32), float32(-math.MaxFloat32)
	check.Equal(t, safe.Subtract(pf32, nf32), pf32, "float32 positive")
	check.Equal(t, safe.Subtract(nf32, pf32), nf32, "float32 negative")

	pf64, nf64 := float64(math.MaxFloat64), float64(-math.MaxFloat64)
	check.Equal(t, safe.Subtract(pf64, nf64), pf64, "float64 positive")
	check.Equal(t, safe.Subtract(nf64, pf64), nf64, "float64 negative")
}

var globalSubtractUint8 uint8

func BenchmarkSubtractUint8(b *testing.B) {
	var x, y, z uint8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		y = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		b.StartTimer()
		z = x - y
	}
	globalSubtractUint8 = z
}

func BenchmarkSafeSubtractUint8(b *testing.B) {
	var x, y, z uint8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		y = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		b.StartTimer()
		z = safe.Subtract(x, y)
	}
	globalSubtractUint8 = z
}

var globalSubtractInt8 int8

func BenchmarkSubtractInt8(b *testing.B) {
	var x, y, z int8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		y = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		b.StartTimer()
		z = x - y
	}
	globalSubtractInt8 = z
}

func BenchmarkSafeSubtractInt8(b *testing.B) {
	var x, y, z int8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		y = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		b.StartTimer()
		z = safe.Subtract(x, y)
	}
	globalSubtractInt8 = z
}

var globalSubtractUint16 uint16

func BenchmarkSubtractUint16(b *testing.B) {
	var x, y, z uint16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		y = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		b.StartTimer()
		z = x - y
	}
	globalSubtractUint16 = z
}

func BenchmarkSafeSubtractUint16(b *testing.B) {
	var x, y, z uint16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		y = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		b.StartTimer()
		z = safe.Subtract(x, y)
	}
	globalSubtractUint16 = z
}

var globalSubtractInt16 int16

func BenchmarkSubtractInt16(b *testing.B) {
	var x, y, z int16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		y = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		b.StartTimer()
		z = x - y
	}
	globalSubtractInt16 = z
}

func BenchmarkSafeSubtractInt16(b *testing.B) {
	var x, y, z int16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		y = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		b.StartTimer()
		z = safe.Subtract(x, y)
	}
	globalSubtractInt16 = z
}

var globalSubtractUint32 uint32

func BenchmarkSubtractUint32(b *testing.B) {
	var x, y, z uint32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint32()
		y = rand.Uint32()
		b.StartTimer()
		z = x - y
	}
	globalSubtractUint32 = z
}

func BenchmarkSafeSubtractUint32(b *testing.B) {
	var x, y, z uint32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint32()
		y = rand.Uint32()
		b.StartTimer()
		z = safe.Subtract(x, y)
	}
	globalSubtractUint32 = z
}

var globalSubtractInt32 int32

func BenchmarkSubtractInt32(b *testing.B) {
	var x, y, z int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int32()
		y = rand.Int32()
		b.StartTimer()
		z = x - y
	}
	globalSubtractInt32 = z
}

func BenchmarkSafeSubtractInt32(b *testing.B) {
	var x, y, z int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int32()
		y = rand.Int32()
		b.StartTimer()
		z = safe.Subtract(x, y)
	}
	globalSubtractInt32 = z
}

var globalSubtractFloat32 float32

func BenchmarkSubtractFloat32(b *testing.B) {
	var x, y, z float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float32()
		y = rand.Float32()
		b.StartTimer()
		z = x - y
	}
	globalSubtractFloat32 = z
}

func BenchmarkSafeSubtractFloat32(b *testing.B) {
	var x, y, z float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float32()
		y = rand.Float32()
		b.StartTimer()
		z = safe.Subtract(x, y)
	}
	globalSubtractFloat32 = z
}

var globalSubtractUint64 uint64

func BenchmarkSubtractUint64(b *testing.B) {
	var x, y, z uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint64()
		y = rand.Uint64()
		b.StartTimer()
		z = x - y
	}
	globalSubtractUint64 = z
}

func BenchmarkSafeSubtractUint64(b *testing.B) {
	var x, y, z uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint64()
		y = rand.Uint64()
		b.StartTimer()
		z = safe.Subtract(x, y)
	}
	globalSubtractUint64 = z
}

var globalSubtractInt64 int64

func BenchmarkSubtractInt64(b *testing.B) {
	var x, y, z int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int64()
		y = rand.Int64()
		b.StartTimer()
		z = x - y
	}
	globalSubtractInt64 = z
}

func BenchmarkSafeSubtractInt64(b *testing.B) {
	var x, y, z int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int64()
		y = rand.Int64()
		b.StartTimer()
		z = safe.Subtract(x, y)
	}
	globalSubtractInt64 = z
}

var globalSubtractFloat64 float64

func BenchmarkSubtractFloat64(b *testing.B) {
	var x, y, z float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float64()
		y = rand.Float64()
		b.StartTimer()
		z = x - y
	}
	globalSubtractFloat64 = z
}

func BenchmarkSafeSubtractFloat64(b *testing.B) {
	var x, y, z float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float64()
		y = rand.Float64()
		b.StartTimer()
		z = safe.Subtract(x, y)
	}
	globalSubtractFloat64 = z
}
