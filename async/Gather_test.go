package async_test

import (
	"github.com/joaonrb/go-lib/async"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGatherShouldReturnOnlyWhenAllFuturesAreFinished(t *testing.T) {
	c := make(chan int)
	number := func(v int) func() int { return func() int { return v } }
	futures := []async.Future[int]{
		async.Run(number(1)),
		async.Run(number(2)),
		async.Run(number(3)),
		async.Run(number(4)),
		async.Run(number(5)),
		async.Run(func() int {
			return <-c
		}),
	}
	future := async.Gather(futures...)
	assert.Equal(t, async.StatusWorking, future.Status())
	c <- 6
	future.Wait()
	assert.Equal(t, async.StatusFinished, future.Status())
	for i, value := range future.Value() {
		assert.Equal(t, i+1, value.Value())
	}
}
