package safe_test

import (
	"math/rand/v2"
	"testing"

	"github.com/samborkent/check"
	"github.com/samborkent/safe"
)

func TestChannelClose(t *testing.T) {
	t.Parallel()

	ch := safe.NewChannel[int](0)
	ch.Close()
	ch.Close()
}

func TestChannelPushPop(t *testing.T) {
	t.Parallel()

	for range 1000000 {
		ch := safe.NewChannel[int](0)
		random1 := rand.Int()
		go ch.Push(random1)
		check.Equal(t, ch.Pop(), random1, "before close")

		ch.Close()
		random2 := rand.Int()
		go ch.Push(random2)
		check.Equal(t, ch.Pop(), 0, "after close")
	}
}
