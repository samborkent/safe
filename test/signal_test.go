package safe_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/samborkent/safe"
)

var innerText *string

func TestSignal(t *testing.T) {
	counter := safe.NewSignal(0)

	isEven := (counter.Get() & 1) == 0
	parity := safe.Ternary(&isEven, "even", "odd")
	innerText = &parity

	fmt.Println(innerText)

	ticker := time.NewTicker(time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			counter += 1

			fmt.Println(innerText)
		}
	}
}
