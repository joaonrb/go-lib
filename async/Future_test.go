package async_test

import (
	"github.com/joaonrb/go-lib/async"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFutureValueShouldPanicIfFutureIsNotStarted(t *testing.T) {
	var future async.Future[any]
	assert.Panics(t, func() {
		future.Value()
	})
}

func TestFutureMustValueShouldPanicIfFutureIsNotStarted(t *testing.T) {
	var future async.Future[any]
	assert.Panics(t, func() {
		future.MustValue()
	})
}

func TestFutureChannelShouldPanicIfFutureIsNotStarted(t *testing.T) {
	var future async.Future[any]
	assert.Panics(t, func() {
		future.Channel()
	})
}

func TestFutureErrorChannelShouldPanicIfFutureIsNotStarted(t *testing.T) {
	var future async.Future[any]
	assert.Panics(t, func() {
		future.ErrorChannel()
	})
}

func TestFutureStringShouldReturnCorrectStringWhenFutureIsNotStarted(t *testing.T) {
	var future async.Future[int]
	assert.Equal(t, "Future[int]{Status: NotStarted}", future.String())
}

func TestFutureStringShouldReturnCorrectStringWhenFutureIsInProgress(t *testing.T) {
	future := async.Run(func() int {
		time.Sleep(time.Second)
		return 0
	})
	assert.Equal(t, "Future[int]{Status: Working}", future.String())
}

func TestFutureStringShouldReturnCorrectStringWhenFutureIsFinished(t *testing.T) {
	future := async.Run(func() int {
		return 99
	})
	future.Wait()
	assert.Equal(t, "Future[int]{Status: Finished, Value: 99}", future.String())
}
