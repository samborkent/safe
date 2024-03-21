package safe_test

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/samborkent/safe"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	var au8, bu8 uint8 = 0, math.MaxUint8
	assert.Equal(t, uint8(0), safe.Subtract(au8, bu8))

	var au16, bu16 uint16 = 0, math.MaxUint16
	assert.Equal(t, uint16(0), safe.Subtract(au16, bu16))

	var au32, bu32 uint32 = 0, math.MaxUint32
	assert.Equal(t, uint32(0), safe.Subtract(au32, bu32))

	var au64, bu64 uint64 = 0, math.MaxUint64
	assert.Equal(t, uint64(0), safe.Subtract(au64, bu64))

	var ap8, bp8 int8 = math.MinInt8, math.MaxInt8
	assert.Equal(t, int8(math.MinInt8), safe.Subtract(ap8, bp8))

	var an8, bn8 int8 = math.MaxInt8, math.MinInt8
	assert.Equal(t, int8(math.MaxInt8), safe.Subtract(an8, bn8))

	var ap16, bp16 int16 = math.MinInt16, math.MaxInt16
	assert.Equal(t, int16(math.MinInt16), safe.Subtract(ap16, bp16))

	var an16, bn16 int16 = math.MaxInt16, math.MinInt16
	assert.Equal(t, int16(math.MaxInt16), safe.Subtract(an16, bn16))

	var ap32, bp32 int32 = math.MinInt32, math.MaxInt32
	assert.Equal(t, int32(math.MinInt32), safe.Subtract(ap32, bp32))

	var an32, bn32 int32 = math.MaxInt32, math.MinInt32
	assert.Equal(t, int32(math.MaxInt32), safe.Subtract(an32, bn32))

	var ap64, bp64 int64 = math.MinInt64, math.MaxInt64
	assert.Equal(t, int64(math.MinInt64), safe.Subtract(ap64, bp64))

	var an64, bn64 int64 = math.MaxInt64, math.MinInt64
	assert.Equal(t, int64(math.MaxInt64), safe.Subtract(an64, bn64))

	var apf32, bpf32 float32 = -math.MaxFloat32, math.MaxFloat32
	assert.Equal(t, float32(-math.MaxFloat32), safe.Subtract(apf32, bpf32))

	var anf32, bnf32 float32 = math.MaxFloat32, -math.MaxFloat32
	assert.Equal(t, float32(math.MaxFloat32), safe.Subtract(anf32, bnf32))

	var apf64, bpf64 float64 = -math.MaxFloat64, math.MaxFloat64
	assert.Equal(t, float64(-math.MaxFloat64), safe.Subtract(apf64, bpf64))

	var anf64, bnf64 float64 = math.MaxFloat64, -math.MaxFloat64
	assert.Equal(t, float64(math.MaxFloat64), safe.Subtract(anf64, bnf64))
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
