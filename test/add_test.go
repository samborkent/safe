package safe_test

import (
	"math"
	"math/rand/v2"
	"testing"

	"safe"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	var au8, bu8 uint8 = math.MaxUint8, math.MaxUint8
	assert.Equal(t, uint8(math.MaxUint8), safe.Add(au8, bu8))

	var au16, bu16 uint16 = math.MaxUint16, math.MaxUint16
	assert.Equal(t, uint16(math.MaxUint16), safe.Add(au16, bu16))

	var au32, bu32 uint32 = math.MaxUint32, math.MaxUint32
	assert.Equal(t, uint32(math.MaxUint32), safe.Add(au32, bu32))

	var au64, bu64 uint64 = math.MaxUint64, math.MaxUint64
	assert.Equal(t, uint64(math.MaxUint64), safe.Add(au64, bu64))

	var ap8, bp8 int8 = math.MaxInt8, math.MaxInt8
	assert.Equal(t, int8(math.MaxInt8), safe.Add(ap8, bp8))

	var an8, bn8 int8 = math.MinInt8, math.MinInt8
	assert.Equal(t, int8(math.MinInt8), safe.Add(an8, bn8))

	var ap16, bp16 int16 = math.MaxInt16, math.MaxInt16
	assert.Equal(t, int16(math.MaxInt16), safe.Add(ap16, bp16))

	var an16, bn16 int16 = math.MinInt16, math.MinInt16
	assert.Equal(t, int16(math.MinInt16), safe.Add(an16, bn16))

	var ap32, bp32 int32 = math.MaxInt32, math.MaxInt32
	assert.Equal(t, int32(math.MaxInt32), safe.Add(ap32, bp32))

	var an32, bn32 int32 = math.MinInt32, math.MinInt32
	assert.Equal(t, int32(math.MinInt32), safe.Add(an32, bn32))

	var ap64, bp64 int64 = math.MaxInt64, math.MaxInt64
	assert.Equal(t, int64(math.MaxInt64), safe.Add(ap64, bp64))

	var an64, bn64 int64 = math.MinInt64, math.MinInt64
	assert.Equal(t, int64(math.MinInt64), safe.Add(an64, bn64))

	var apf32, bpf32 float32 = math.MaxFloat32, math.MaxFloat32
	assert.Equal(t, float32(math.MaxFloat32), safe.Add(apf32, bpf32))

	var anf32, bnf32 float32 = -math.MaxFloat32, -math.MaxFloat32
	assert.Equal(t, float32(-math.MaxFloat32), safe.Add(anf32, bnf32))

	var apf64, bpf64 float64 = math.MaxFloat64, math.MaxFloat64
	assert.Equal(t, float64(math.MaxFloat64), safe.Add(apf64, bpf64))

	var anf64, bnf64 float64 = -math.MaxFloat64, -math.MaxFloat64
	assert.Equal(t, float64(-math.MaxFloat64), safe.Add(anf64, bnf64))

	var api, bpi int = math.MaxInt, math.MaxInt
	assert.Equal(t, int(math.MaxInt), safe.Add(api, bpi))

	var ani, bni int = math.MinInt, math.MinInt
	assert.Equal(t, int(math.MinInt), safe.Add(ani, bni))
}

var globalAddUint8 uint8

func BenchmarkAddUint8(b *testing.B) {
	var x, y, z uint8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		y = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		b.StartTimer()
		z = x + y
	}
	globalAddUint8 = z
}

func BenchmarkSafeAddUint8(b *testing.B) {
	var x, y, z uint8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		y = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddUint8 = z
}

var globalAddInt8 int8

func BenchmarkAddInt8(b *testing.B) {
	var x, y, z int8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		y = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		b.StartTimer()
		z = x + y
	}
	globalAddInt8 = z
}

func BenchmarkSafeAddInt8(b *testing.B) {
	var x, y, z int8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		y = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddInt8 = z
}

var globalAddUint16 uint16

func BenchmarkAddUint16(b *testing.B) {
	var x, y, z uint16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		y = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		b.StartTimer()
		z = x + y
	}
	globalAddUint16 = z
}

func BenchmarkSafeAddUint16(b *testing.B) {
	var x, y, z uint16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		y = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddUint16 = z
}

var globalAddInt16 int16

func BenchmarkAddInt16(b *testing.B) {
	var x, y, z int16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		y = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		b.StartTimer()
		z = x + y
	}
	globalAddInt16 = z
}

func BenchmarkSafeAddInt16(b *testing.B) {
	var x, y, z int16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		y = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddInt16 = z
}

var globalAddUint32 uint32

func BenchmarkAddUint32(b *testing.B) {
	var x, y, z uint32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint32()
		y = rand.Uint32()
		b.StartTimer()
		z = x + y
	}
	globalAddUint32 = z
}

func BenchmarkSafeAddUint32(b *testing.B) {
	var x, y, z uint32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint32()
		y = rand.Uint32()
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddUint32 = z
}

var globalAddInt32 int32

func BenchmarkAddInt32(b *testing.B) {
	var x, y, z int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int32()
		y = rand.Int32()
		b.StartTimer()
		z = x + y
	}
	globalAddInt32 = z
}

func BenchmarkSafeAddInt32(b *testing.B) {
	var x, y, z int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int32()
		y = rand.Int32()
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddInt32 = z
}

var globalAddFloat32 float32

func BenchmarkAddFloat32(b *testing.B) {
	var x, y, z float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float32()
		y = rand.Float32()
		b.StartTimer()
		z = x + y
	}
	globalAddFloat32 = z
}

func BenchmarkSafeAddFloat32(b *testing.B) {
	var x, y, z float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float32()
		y = rand.Float32()
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddFloat32 = z
}

var globalAddUint64 uint64

func BenchmarkAddUint64(b *testing.B) {
	var x, y, z uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint64()
		y = rand.Uint64()
		b.StartTimer()
		z = x + y
	}
	globalAddUint64 = z
}

func BenchmarkSafeAddUint64(b *testing.B) {
	var x, y, z uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint64()
		y = rand.Uint64()
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddUint64 = z
}

var globalAddInt64 int64

func BenchmarkAddInt64(b *testing.B) {
	var x, y, z int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int64()
		y = rand.Int64()
		b.StartTimer()
		z = x + y
	}
	globalAddInt64 = z
}

func BenchmarkSafeAddInt64(b *testing.B) {
	var x, y, z int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int64()
		y = rand.Int64()
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddInt64 = z
}

var globalAddFloat64 float64

func BenchmarkAddFloat64(b *testing.B) {
	var x, y, z float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float64()
		y = rand.Float64()
		b.StartTimer()
		z = x + y
	}
	globalAddFloat64 = z
}

func BenchmarkSafeAddFloat64(b *testing.B) {
	var x, y, z float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float64()
		y = rand.Float64()
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddFloat64 = z
}

var globalAddInt int

func BenchmarkAddInt(b *testing.B) {
	var x, y, z int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int()
		y = rand.Int()
		b.StartTimer()
		z = x + y
	}
	globalAddInt = z
}

func BenchmarkSafeAddInt(b *testing.B) {
	var x, y, z int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int()
		y = rand.Int()
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddInt = z
}

var globalAddUint uint

func BenchmarkAddUint(b *testing.B) {
	var x, y, z uint
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint(rand.Uint64())
		y = uint(rand.Uint64())
		b.StartTimer()
		z = x + y
	}
	globalAddUint = z
}

func BenchmarkSafeAddUint(b *testing.B) {
	var x, y, z uint
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint(rand.Uint64())
		y = uint(rand.Uint64())
		b.StartTimer()
		z = safe.Add(x, y)
	}
	globalAddUint = z
}
