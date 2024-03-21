package safe_test

import (
	"testing"

	"github.com/samborkent/safe"

	"github.com/stretchr/testify/assert"
)

func TestChannelClose(t *testing.T) {
	ch := safe.NewChannel[int](0)
	ch.Close()
	ch.Close()
}

func TestChannelPushPop(t *testing.T) {
	for range 1000000 {
		ch := safe.NewChannel[int](0)
		go ch.Push(5)
		assert.Equal(t, 5, ch.Pop())
		ch.Close()
		go ch.Push(6)
		assert.Equal(t, 0, ch.Pop())
	}
}
