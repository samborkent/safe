package safe_test

import (
	"math"
	"math/rand/v2"
	"testing"

	"github.com/samborkent/check"
	"github.com/samborkent/safe"
)

func TestDivide(t *testing.T) {
	t.Parallel()

	check.Equal(t, safe.Divide(0, 1), 0, "zero numerator")

	var au8, bu8 uint8 = 1, 0
	check.Equal(t, safe.Divide(au8, bu8), uint8(math.MaxUint8), "uint8")

	var au16, bu16 uint16 = 1, 0
	check.Equal(t, safe.Divide(au16, bu16), uint16(math.MaxUint16), "uint16")

	var au32, bu32 uint32 = 1, 0
	check.Equal(t, safe.Divide(au32, bu32), uint32(math.MaxUint32), "uint32")

	var au64, bu64 uint64 = 1, 0
	check.Equal(t, safe.Divide(au64, bu64), uint64(math.MaxUint64), "uint64")

	var au, bu uint = 1, 0
	check.Equal(t, safe.Divide(au, bu), uint(math.MaxUint), "uint")

	var ap8, an8, bi8 int8 = 1, -1, 0
	check.Equal(t, safe.Divide(ap8, bi8), int8(math.MaxInt8), "int8 positive")
	check.Equal(t, safe.Divide(an8, bi8), int8(math.MinInt8), "int8 negative")

	var ap16, an16, bi16 int16 = 1, -1, 0
	check.Equal(t, safe.Divide(ap16, bi16), int16(math.MaxInt16), "int16 positive")
	check.Equal(t, safe.Divide(an16, bi16), int16(math.MinInt16), "int16 negative")

	var ap32, an32, bi32 int32 = 1, -1, 0
	check.Equal(t, safe.Divide(ap32, bi32), int32(math.MaxInt32), "int32 positive")
	check.Equal(t, safe.Divide(an32, bi32), int32(math.MinInt32), "int32 negative")

	var ap64, an64, bi64 int64 = 1, -1, 0
	check.Equal(t, safe.Divide(ap64, bi64), int64(math.MaxInt64), "int64 positive")
	check.Equal(t, safe.Divide(an64, bi64), int64(math.MinInt64), "int64 negative")

	var aip, ain, bi int = 1, -1, 0
	check.Equal(t, safe.Divide(aip, bi), int(math.MaxInt), "int64 positive")
	check.Equal(t, safe.Divide(ain, bi), int(math.MinInt), "int64 negative")

	var apf32, anf32, bf32 float32 = 1, -1, 0
	check.Equal(t, safe.Divide(apf32, bf32), float32(math.MaxFloat32), "float32 positive")
	check.Equal(t, safe.Divide(anf32, bf32), float32(-math.MaxFloat32), "float32 negative")

	var apf64, anf64, bf64 float64 = 1, -1, 0
	check.Equal(t, safe.Divide(apf64, bf64), float64(math.MaxFloat64), "float64 positive")
	check.Equal(t, safe.Divide(anf64, bf64), float64(-math.MaxFloat64), "float64 negative")
}

var globalDivideUint8 uint8

func BenchmarkDivideUint8(b *testing.B) {
	var x, y, z uint8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		y = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		if y == 0 {
			y++
		}
		b.StartTimer()
		z = x / y
	}
	globalDivideUint8 = z
}

func BenchmarkSafeDivideUint8(b *testing.B) {
	var x, y, z uint8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		y = uint8(float64(rand.Uint32()) * float64(math.MaxUint8) / float64(math.MaxUint32))
		b.StartTimer()
		z = safe.Divide(x, y)
	}
	globalDivideUint8 = z
}

var globalDivideInt8 int8

func BenchmarkDivideInt8(b *testing.B) {
	var x, y, z int8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		y = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		if y == 0 {
			y++
		}
		b.StartTimer()
		z = x / y
	}
	globalDivideInt8 = z
}

func BenchmarkSafeDivideInt8(b *testing.B) {
	var x, y, z int8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		y = int8(float64(rand.Int32()) * float64(-math.MinInt8) / float64(-math.MinInt32))
		b.StartTimer()
		z = safe.Divide(x, y)
	}
	globalDivideInt8 = z
}

var globalDivideUint16 uint16

func BenchmarkDivideUint16(b *testing.B) {
	var x, y, z uint16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		y = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		if y == 0 {
			y++
		}
		b.StartTimer()
		z = x / y
	}
	globalDivideUint16 = z
}

func BenchmarkSafeDivideUint16(b *testing.B) {
	var x, y, z uint16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		y = uint16(float64(rand.Uint32()) * float64(math.MaxUint16) / float64(math.MaxUint32))
		b.StartTimer()
		z = safe.Divide(x, y)
	}
	globalDivideUint16 = z
}

var globalDivideInt16 int16

func BenchmarkDivideInt16(b *testing.B) {
	var x, y, z int16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		y = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		if y == 0 {
			y++
		}
		b.StartTimer()
		z = x / y
	}
	globalDivideInt16 = z
}

func BenchmarkSafeDivideInt16(b *testing.B) {
	var x, y, z int16
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		y = int16(float64(rand.Int32()) * float64(-math.MinInt16) / float64(-math.MinInt32))
		b.StartTimer()
		z = safe.Divide(x, y)
	}
	globalDivideInt16 = z
}

var globalDivideUint32 uint32

func BenchmarkDivideUint32(b *testing.B) {
	var x, y, z uint32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint32()
		y = rand.Uint32()
		if y == 0 {
			y++
		}
		b.StartTimer()
		z = x / y
	}
	globalDivideUint32 = z
}

func BenchmarkSafeDivideUint32(b *testing.B) {
	var x, y, z uint32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint32()
		y = rand.Uint32()
		b.StartTimer()
		z = safe.Divide(x, y)
	}
	globalDivideUint32 = z
}

var globalDivideInt32 int32

func BenchmarkDivideInt32(b *testing.B) {
	var x, y, z int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int32()
		y = rand.Int32()
		if y == 0 {
			y++
		}
		b.StartTimer()
		z = x / y
	}
	globalDivideInt32 = z
}

func BenchmarkSafeDivideInt32(b *testing.B) {
	var x, y, z int32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int32()
		y = rand.Int32()
		b.StartTimer()
		z = safe.Divide(x, y)
	}
	globalDivideInt32 = z
}

var globalDivideFloat32 float32

func BenchmarkDivideFloat32(b *testing.B) {
	var x, y, z float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float32()
		y = rand.Float32()
		if y == 0 {
			y++
		}
		b.StartTimer()
		z = x / y
	}
	globalDivideFloat32 = z
}

func BenchmarkSafeDivideFloat32(b *testing.B) {
	var x, y, z float32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float32()
		y = rand.Float32()
		b.StartTimer()
		z = safe.Divide(x, y)
	}
	globalDivideFloat32 = z
}

var globalDivideUint64 uint64

func BenchmarkDivideUint64(b *testing.B) {
	var x, y, z uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint64()
		y = rand.Uint64()
		if y == 0 {
			y++
		}
		b.StartTimer()
		z = x / y
	}
	globalDivideUint64 = z
}

func BenchmarkSafeDivideUint64(b *testing.B) {
	var x, y, z uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Uint64()
		y = rand.Uint64()
		b.StartTimer()
		z = safe.Divide(x, y)
	}
	globalDivideUint64 = z
}

var globalDivideInt64 int64

func BenchmarkDivideInt64(b *testing.B) {
	var x, y, z int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int64()
		y = rand.Int64()
		if y == 0 {
			y++
		}
		b.StartTimer()
		z = x / y
	}
	globalDivideInt64 = z
}

func BenchmarkSafeDivideInt64(b *testing.B) {
	var x, y, z int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Int64()
		y = rand.Int64()
		b.StartTimer()
		z = safe.Divide(x, y)
	}
	globalDivideInt64 = z
}

var globalDivideFloat64 float64

func BenchmarkDivideFloat64(b *testing.B) {
	var x, y, z float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float64()
		y = rand.Float64()
		if y == 0 {
			y++
		}
		b.StartTimer()
		z = x / y
	}
	globalDivideFloat64 = z
}

func BenchmarkSafeDivideFloat64(b *testing.B) {
	var x, y, z float64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		x = rand.Float64()
		y = rand.Float64()
		b.StartTimer()
		z = safe.Divide(x, y)
	}
	globalDivideFloat64 = z
}
