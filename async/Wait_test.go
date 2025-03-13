package async_test

import (
	"testing"

	"github.com/joaonrb/go-lib/async"
	"github.com/stretchr/testify/assert"
)

func TestWaitShouldReturnFuturesAsTheyFinish(t *testing.T) {
	number := func(v int) (func() int, chan int) {
		c := make(chan int)
		return func() int {
			<-c
			return v
		}, c
	}
	tries := 10
	futures := make([]async.Future[int], tries)
	channels := make([]chan int, tries)
	for i := 0; i < tries; i++ {
		var f func() int
		f, channels[i] = number(i)
		futures[i] = async.Run(f)
	}
	future := async.Wait[int](futures...)
	channel := future.Value()
	assertWaitReturnsN(t, 3, channels, channel)
	assertWaitReturnsN(t, 0, channels, channel)
	assertWaitReturnsN(t, 8, channels, channel)
	assertWaitReturnsN(t, 4, channels, channel)
	assertWaitReturnsN(t, 2, channels, channel)
	assertWaitReturnsN(t, 1, channels, channel)
	assertWaitReturnsN(t, 6, channels, channel)
	assertWaitReturnsN(t, 7, channels, channel)
	assertWaitReturnsN(t, 5, channels, channel)
	assertWaitReturnsN(t, 9, channels, channel)

}

func assertWaitReturnsN(
	t *testing.T,
	n int,
	channels []chan int,
	channel <-chan async.Future[int],
) {
	go func() { channels[n] <- 0 }()
	assert.Equal(t, n, (<-channel).MustValue().TryValue())
}
