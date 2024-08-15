package safe_test

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/samborkent/safe"
	"github.com/samborkent/safe/thelper"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	var au8, bu8 uint8 = math.MaxUint8, math.MaxUint8
	thelper.Equal(t, safe.Add(au8, bu8), uint8(math.MaxUint8), "uint8")

	var au16, bu16 uint16 = math.MaxUint16, math.MaxUint16
	thelper.Equal(t, safe.Add(au16, bu16), uint16(math.MaxUint16), "uint16")

	var au32, bu32 uint32 = math.MaxUint32, math.MaxUint32
	thelper.Equal(t, safe.Add(au32, bu32), uint32(math.MaxUint32), "uint32")

	var au64, bu64 uint64 = math.MaxUint64, math.MaxUint64
	thelper.Equal(t, safe.Add(au64, bu64), uint64(math.MaxUint64), "uint64")

	var ap8, bp8 int8 = math.MaxInt8, math.MaxInt8
	thelper.Equal(t, safe.Add(ap8, bp8), int8(math.MaxInt8), "int8 overflow")

	var an8, bn8 int8 = math.MinInt8, math.MinInt8
	thelper.Equal(t, safe.Add(an8, bn8), int8(math.MinInt8), "int8 underflow")

	var ap16, bp16 int16 = math.MaxInt16, math.MaxInt16
	thelper.Equal(t, safe.Add(ap16, bp16), int16(math.MaxInt16), "int16 overflow")

	var an16, bn16 int16 = math.MinInt16, math.MinInt16
	thelper.Equal(t, safe.Add(an16, bn16), int16(math.MinInt16), "int16 underflow")

	var ap32, bp32 int32 = math.MaxInt32, math.MaxInt32
	thelper.Equal(t, safe.Add(ap32, bp32), int32(math.MaxInt32), "int32 overflow")

	var an32, bn32 int32 = math.MinInt32, math.MinInt32
	thelper.Equal(t, safe.Add(an32, bn32), int32(math.MinInt32), "int32 underflow")

	var ap64, bp64 int64 = math.MaxInt64, math.MaxInt64
	thelper.Equal(t, safe.Add(ap64, bp64), int64(math.MaxInt64), "int64 overflow")

	var an64, bn64 int64 = math.MinInt64, math.MinInt64
	thelper.Equal(t, safe.Add(an64, bn64), int64(math.MinInt64), "int64 underflow")

	var api, bpi int = math.MaxInt, math.MaxInt
	thelper.Equal(t, safe.Add(api, bpi), int(math.MaxInt), "int overflow")

	var ani, bni int = math.MinInt, math.MinInt
	thelper.Equal(t, safe.Add(ani, bni), int(math.MinInt), "int underflow")

	var apf32, bpf32 float32 = math.MaxFloat32, math.MaxFloat32
	thelper.Equal(t, safe.Add(apf32, bpf32), float32(math.MaxFloat32), "float32 overflow")

	var anf32, bnf32 float32 = -math.MaxFloat32, -math.MaxFloat32
	thelper.Equal(t, safe.Add(anf32, bnf32), float32(-math.MaxFloat32), "float32 underflow")

	var apf64, bpf64 float64 = math.MaxFloat64, math.MaxFloat64
	thelper.Equal(t, safe.Add(apf64, bpf64), float64(math.MaxFloat64), "float64 overflow")

	var anf64, bnf64 float64 = -math.MaxFloat64, -math.MaxFloat64
	thelper.Equal(t, safe.Add(anf64, bnf64), float64(-math.MaxFloat64), "float64 underflow")
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
