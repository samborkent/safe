package safe_test

import (
	"fmt"
	"testing"

	"github.com/samborkent/safe"
	"github.com/samborkent/safe/thelper"
)

func TestMapRange(t *testing.T) {
	m := safe.NewMap[string, float64](0)

	key := []string{"hello", "foo", "bar"}
	val := []float64{0.6, 0.3, 0.2}

	m.Store(key[0], val[0])
	m.Store(key[1], val[1])
	m.Store(key[2], val[2])

	index := 0

	for k, v := range m.Range() {
		thelper.Equal(t, k, key[index], fmt.Sprintf("key %d", index))
		thelper.Equal(t, v, val[index], fmt.Sprintf("val %d", index))
		index++
	}
}
