package safe

import (
	"testing"

	"github.com/samborkent/check"
	"github.com/samborkent/safe"
)

// TODO: add checks for struct -> interface assertion

type testStruct struct {
	value int
}

func (testStruct) Test() {}

type testStruct2 struct {
	value int
}

func (testStruct2) Test() {}

// Explicitely ignore unused warning
var _ = testStruct2{value: 0}

type testStruct3 struct{}

type testInterface interface {
	Test()
}

type testInterface2 interface {
	Test()
}

type testInterface3 interface {
	OtherTest()
}

func TestTypeAssert(t *testing.T) {
	// Pass: int
	a1 := 10
	a2 := safe.TypeAssert[int](a1)
	check.Equal(t, a2, a1, "int")

	// // Fail: int -> string
	// a3 := 10
	// a4 := safe.TypeAssert[string](a3)
	// assert.NotEqualValues(t, a3, a4)

	// Pass: string
	b1 := "string"
	b2 := safe.TypeAssert[string](b1)
	check.Equal(t, b1, b2)

	// // Fail: string -> int
	// b3 := "string"
	// b4 := safe.TypeAssert[int](b3)
	// assert.NotEqualValues(t, b3, b4)

	// Pass: struct
	c1 := testStruct{
		value: 123,
	}
	c2 := safe.TypeAssert[testStruct](c1)
	check.Equal(t, c1, c2)

	// // Fail: struct -> different struct with same values
	// c3 := testStruct{
	// 	value: 123,
	// }
	// c4 := safe.TypeAssert[testStruct2](c3)
	// assert.NotEqualValues(t, c3, c4)

	// // Fail: struct without values -> different struct without values
	// c5 := struct{}{}
	// c6 := safe.TypeAssert[testStruct3](c5)
	// assert.NotEqual(t, c5, c6)

	// Pass: struct pointer
	d1 := &testStruct{
		value: 123,
	}
	d2 := safe.TypeAssert[*testStruct](d1)
	check.Equal(t, d1, d2)

	// // Fail: struct pointer -> different struct pointer with same values
	// d3 := &testStruct{
	// 	value: 123,
	// }
	// d4 := safe.TypeAssert[*testStruct2](d3)
	// assert.NotEqualValues(t, d3, d4)

	// // Fail: struct pointer without values -> different struct pointer without values
	// d5 := &struct{}{}
	// d6 := safe.TypeAssert[*testStruct3](d5)
	// assert.NotEqualValues(t, d5, d6)

	// Pass: struct which implements interface
	var e1 testStruct
	e2 := safe.TypeAssert[testInterface](e1)
	check.Equal(t, testInterface(e1), e2)

	// Fail: struct which does not implement interface
	var e3 testStruct
	e4 := safe.TypeAssert[testInterface3](e3)
	check.Nil(t, e4)

	// Pass: pointer to struct to interface
	var f1 testStruct
	f2 := safe.TypeAssert[testInterface](&f1)
	check.Equal(t, testInterface(&f1), f2)

	// Fail: interface -> different interface
	var f3 testStruct
	f4 := safe.TypeAssert[testInterface3](&f3)
	check.Nil(t, f4)

	// Pass: nil interface
	var g1 testInterface
	g2 := safe.TypeAssert[testInterface2](g1)
	check.Equal(t, testInterface2(g1), g2)
}

// // TODO: fix
// func TestRequireTypeAssert(t *testing.T) {
// 	// Pass: int
// 	a1 := 10
// 	a2, err := safe.RequireTypeAssert[int](a1)
// 	assert.check.Equal(t, a1, a2)
// 	assert.Nil(t, err)

// 	// Fail: int -> string
// 	a3 := 10
// 	a4, err := safe.RequireTypeAssert[string](a3)
// 	assert.NotEqualValues(t, a3, a4)
// 	assert.NotNil(t, err)

// 	// Pass: string
// 	b1 := "string"
// 	b2, err := safe.RequireTypeAssert[string](b1)
// 	assert.check.Equal(t, b1, b2)
// 	assert.Nil(t, err)

// 	// Fail: string -> int
// 	b3 := "string"
// 	b4, err := safe.RequireTypeAssert[int](b3)
// 	assert.NotEqualValues(t, b3, b4)
// 	assert.NotNil(t, err)

// 	// Pass: struct
// 	c1 := testStruct{
// 		value: 123,
// 	}
// 	c2, err := safe.RequireTypeAssert[testStruct](c1)
// 	assert.check.Equal(t, c1, c2)
// 	assert.Nil(t, err)

// 	// Fail: struct -> different struct with same values
// 	c3 := testStruct{
// 		value: 123,
// 	}
// 	c4, err := safe.RequireTypeAssert[testStruct2](c3)
// 	assert.NotEqualValues(t, c3, c4)
// 	assert.NotNil(t, err)

// 	// Fail: struct without values -> different struct without values
// 	c5 := struct{}{}
// 	c6, err := safe.RequireTypeAssert[testStruct3](c5)
// 	assert.Notcheck.Equal(t, c5, c6)
// 	assert.NotNil(t, err)

// 	// Pass: struct pointer
// 	d1 := &testStruct{
// 		value: 123,
// 	}
// 	d2, err := safe.RequireTypeAssert[*testStruct](d1)
// 	assert.check.Equal(t, d1, d2)
// 	assert.Nil(t, err)

// 	// Fail: struct pointer -> different struct pointer with same values
// 	d3 := &testStruct{
// 		value: 123,
// 	}
// 	d4, err := safe.RequireTypeAssert[*testStruct2](d3)
// 	assert.NotEqualValues(t, d3, d4)
// 	assert.NotNil(t, err)

// 	// Fail: struct pointer without values -> different struct pointer without values
// 	d5 := &struct{}{}
// 	d6, err := safe.RequireTypeAssert[*testStruct3](d5)
// 	assert.NotEqualValues(t, d5, d6)
// 	assert.NotNil(t, err)

// 	// Pass: struct which implements interface
// 	var e1 testStruct
// 	e2, err := safe.RequireTypeAssert[testInterface](e1)
// 	assert.check.Equal(t, e1, e2)
// 	assert.Nil(t, err)

// 	// Fail: struct which implements interface -> different interface
// 	var e3 testStruct
// 	e4, err := safe.RequireTypeAssert[testInterface3](e3)
// 	assert.Nil(t, e4)
// 	assert.NotNil(t, err)

// 	// Fail: any -> empty interface
// 	var e5 any
// 	e6, err := safe.RequireTypeAssert[testInterface3](e5)
// 	assert.Nil(t, e6)
// 	assert.NotNil(t, err)

// 	// Pass: interface pointer
// 	var f1 *testInterface
// 	f2, err := safe.RequireTypeAssert[*testInterface](any(f1))
// 	assert.check.Equal(t, f1, f2)
// 	assert.Nil(t, err)

// 	// Fail: interface pointer -> different interface pointer
// 	var f3 *testInterface
// 	f4, err := safe.RequireTypeAssert[*testInterface2](any(f3))
// 	assert.Nil(t, f4)
// 	assert.NotNil(t, err)

// 	// Fail: any pointer -> empty interface pointer
// 	var f5 *any
// 	f6, err := safe.RequireTypeAssert[*testInterface3](f5)
// 	assert.Nil(t, f6)
// 	assert.NotNil(t, err)
// }
