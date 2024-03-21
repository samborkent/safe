package safe_test

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/samborkent/safe"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	var au8, bu8 uint8 = math.MaxUint8, math.MaxUint8
	assert.Equal(t, uint8(math.MaxUint8), safe.Multiply(au8, bu8))

	var au16, bu16 uint16 = math.MaxUint16, math.MaxUint16
	assert.Equal(t, uint16(math.MaxUint16), safe.Multiply(au16, bu16))

	var au32, bu32 uint32 = math.MaxUint32, math.MaxUint32
	assert.Equal(t, uint32(math.MaxUint32), safe.Multiply(au32, bu32))

	var au64, bu64 uint64 = math.MaxUint64, math.MaxUint64
	assert.Equal(t, uint64(math.MaxUint64), safe.Multiply(au64, bu64))

	var ap8, bp8 int8 = math.MaxInt8, math.MaxInt8
	assert.Equal(t, int8(math.MaxInt8), safe.Multiply(ap8, bp8))

	var an8, bn8 int8 = math.MinInt8, math.MinInt8
	assert.Equal(t, int8(math.MinInt8), safe.Multiply(an8, bn8))

	var ap16, bp16 int16 = math.MaxInt16, math.MaxInt16
	assert.Equal(t, int16(math.MaxInt16), safe.Multiply(ap16, bp16))

	var an16, bn16 int16 = math.MinInt16, math.MinInt16
	assert.Equal(t, int16(math.MinInt16), safe.Multiply(an16, bn16))

	var ap32, bp32 int32 = math.MaxInt32, math.MaxInt32
	assert.Equal(t, int32(math.MaxInt32), safe.Multiply(ap32, bp32))

	var an32, bn32 int32 = math.MinInt32, math.MinInt32
	assert.Equal(t, int32(math.MinInt32), safe.Multiply(an32, bn32))

	var ap64, bp64 int64 = math.MaxInt64, math.MaxInt64
	assert.Equal(t, int64(math.MaxInt64), safe.Multiply(ap64, bp64))

	var an64, bn64 int64 = math.MinInt64, math.MinInt64
	assert.Equal(t, int64(math.MinInt64), safe.Multiply(an64, bn64))

	var apf32, bpf32 float32 = math.MaxFloat32, math.MaxFloat32
	assert.Equal(t, float32(math.MaxFloat32), safe.Multiply(apf32, bpf32))

	var anf32, bnf32 float32 = -math.MaxFloat32, -math.MaxFloat32
	assert.Equal(t, float32(-math.MaxFloat32), safe.Multiply(anf32, bnf32))

	var apf64, bpf64 float64 = math.MaxFloat64, math.MaxFloat64
	assert.Equal(t, float64(math.MaxFloat64), safe.Multiply(apf64, bpf64))

	var anf64, bnf64 float64 = -math.MaxFloat64, -math.MaxFloat64
	assert.Equal(t, float64(-math.MaxFloat64), safe.Multiply(anf64, bnf64))
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
