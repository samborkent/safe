package safe_test

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/samborkent/safe"
	"github.com/samborkent/safe/thelper"
)

func TestMultiply(t *testing.T) {
	u8 := uint8(math.MaxUint8)
	thelper.Equal(t, safe.Multiply(u8, u8), u8, "uint8")

	u16 := uint16(math.MaxUint16)
	thelper.Equal(t, safe.Multiply(u16, u16), u16, "uint16")

	u32 := uint32(math.MaxUint32)
	thelper.Equal(t, safe.Multiply(u32, u32), u32, "uint32")

	u64 := uint64(math.MaxUint64)
	thelper.Equal(t, safe.Multiply(u64, u64), u64, "uint64")

	u := uint(math.MaxUint)
	thelper.Equal(t, safe.Multiply(u, u), u, "uint")

	p8 := int8(math.MaxInt8)
	thelper.Equal(t, safe.Multiply(p8, p8), p8, "int8 max")

	n8 := int8(math.MinInt8)
	thelper.Equal(t, safe.Multiply(n8, n8), n8, "int8 min")

	p16 := int16(math.MaxInt16)
	thelper.Equal(t, safe.Multiply(p16, p16), p16, "int16 max")

	n16 := int16(math.MinInt16)
	thelper.Equal(t, safe.Multiply(n16, n16), n16, "int16 min")

	p32 := int32(math.MaxInt32)
	thelper.Equal(t, safe.Multiply(p32, p32), p32, "int32 max")

	n32 := int32(math.MinInt32)
	thelper.Equal(t, safe.Multiply(n32, n32), n32, "int32 min")

	p64 := int64(math.MaxInt64)
	thelper.Equal(t, safe.Multiply(p64, p64), p64, "int64 max")

	n64 := int64(math.MinInt64)
	thelper.Equal(t, safe.Multiply(n64, n64), n64, "int64 min")

	pi := int64(math.MaxInt64)
	thelper.Equal(t, safe.Multiply(pi, pi), pi, "int max")

	ni := int64(math.MinInt64)
	thelper.Equal(t, safe.Multiply(ni, ni), ni, "int min")

	pf32 := float32(math.MaxFloat32)
	thelper.Equal(t, safe.Multiply(pf32, pf32), pf32, "float32 max")

	nf32 := float32(-math.MaxFloat32)
	thelper.Equal(t, safe.Multiply(nf32, nf32), nf32, "float32 min")

	pf64 := float64(math.MaxFloat64)
	thelper.Equal(t, safe.Multiply(pf64, pf64), pf64, "float64 max")

	nf64 := float64(-math.MaxFloat64)
	thelper.Equal(t, safe.Multiply(nf64, nf64), nf64, "float64 min")
}

var globalMultiplyUint8 uint8

func BenchmarkMultiplyUint8(b *testing.B) {
	var x, y, z uint8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		y = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		b.StartTimer()
		z = x * y
	}
	globalMultiplyUint8 = z
}

func BenchmarkSafeMultiplyUint8(b *testing.B) {
	var x, y, z uint8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		y = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		b.StartTimer()
		z = safe.Multiply(x, y)
	}
	globalMultiplyUint8 = z
}

var globalMultiplyInt8 int8

func BenchmarkMultiplyInt8(b *testing.B) {
	var x, y, z int8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		y = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		b.StartTimer()
		z = x * y
	}
	globalMultiplyInt8 = z
}

func BenchmarkSafeMultiplyInt8(b *testing.B) {
	var x, y, z int8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		y = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		b.StartTimer()
		z = safe.Multiply(x, y)
	}
	globalMultiplyInt8 = z
}

var globalMultiplyUint16 uint16

func BenchmarkMultiplyUint16(b *testing.B) {
	var x, y, z uint16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		y = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		b.StartTimer()
		z = x * y
	}
	globalMultiplyUint16 = z
}

func BenchmarkSafeMultiplyUint16(b *testing.B) {
	var x, y, z uint16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		y = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		b.StartTimer()
		z = safe.Multiply(x, y)
	}
	globalMultiplyUint16 = z
}

var globalMultiplyInt16 int16

func BenchmarkMultiplyInt16(b *testing.B) {
	var x, y, z int16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		y = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		b.StartTimer()
		z = x * y
	}
	globalMultiplyInt16 = z
}

func BenchmarkSafeMultiplyInt16(b *testing.B) {
	var x, y, z int16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		y = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		b.StartTimer()
		z = safe.Multiply(x, y)
	}
	globalMultiplyInt16 = z
}

var globalMultiplyUint32 uint32

func BenchmarkMultiplyUint32(b *testing.B) {
	var x, y, z uint32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint32()
		y = rand.Uint32()
		b.StartTimer()
		z = x * y
	}
	globalMultiplyUint32 = z
}

func BenchmarkSafeMultiplyUint32(b *testing.B) {
	var x, y, z uint32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint32()
		y = rand.Uint32()
		b.StartTimer()
		z = safe.Multiply(x, y)
	}
	globalMultiplyUint32 = z
}

var globalMultiplyInt32 int32

func BenchmarkMultiplyInt32(b *testing.B) {
	var x, y, z int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int32()
		y = rand.Int32()
		b.StartTimer()
		z = x * y
	}
	globalMultiplyInt32 = z
}

func BenchmarkSafeMultiplyInt32(b *testing.B) {
	var x, y, z int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int32()
		y = rand.Int32()
		b.StartTimer()
		z = safe.Multiply(x, y)
	}
	globalMultiplyInt32 = z
}

var globalMultiplyFloat32 float32

func BenchmarkMultiplyFloat32(b *testing.B) {
	var x, y, z float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float32()
		y = rand.Float32()
		b.StartTimer()
		z = x * y
	}
	globalMultiplyFloat32 = z
}

func BenchmarkSafeMultiplyFloat32(b *testing.B) {
	var x, y, z float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float32()
		y = rand.Float32()
		b.StartTimer()
		z = safe.Multiply(x, y)
	}
	globalMultiplyFloat32 = z
}

var globalMultiplyUint64 uint64

func BenchmarkMultiplyUint64(b *testing.B) {
	var x, y, z uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint64()
		y = rand.Uint64()
		b.StartTimer()
		z = x * y
	}
	globalMultiplyUint64 = z
}

func BenchmarkSafeMultiplyUint64(b *testing.B) {
	var x, y, z uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint64()
		y = rand.Uint64()
		b.StartTimer()
		z = safe.Multiply(x, y)
	}
	globalMultiplyUint64 = z
}

var globalMultiplyInt64 int64

func BenchmarkMultiplyInt64(b *testing.B) {
	var x, y, z int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int64()
		y = rand.Int64()
		b.StartTimer()
		z = x * y
	}
	globalMultiplyInt64 = z
}

func BenchmarkSafeMultiplyInt64(b *testing.B) {
	var x, y, z int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int64()
		y = rand.Int64()
		b.StartTimer()
		z = safe.Multiply(x, y)
	}
	globalMultiplyInt64 = z
}

var globalMultiplyFloat64 float64

func BenchmarkMultiplyFloat64(b *testing.B) {
	var x, y, z float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float64()
		y = rand.Float64()
		b.StartTimer()
		z = x * y
	}
	globalMultiplyFloat64 = z
}

func BenchmarkSafeMultiplyFloat64(b *testing.B) {
	var x, y, z float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float64()
		y = rand.Float64()
		b.StartTimer()
		z = safe.Multiply(x, y)
	}
	globalMultiplyFloat64 = z
}
