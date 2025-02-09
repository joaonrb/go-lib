package async_test

import (
	"testing"
	"time"

	"github.com/joaonrb/go-lib/async"
	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunShouldReleaseRightAwayAfterBeingCalled(t *testing.T) {
	future := async.Run(func() int {
		time.Sleep(time.Hour)
		return 0
	})
	assert.Equal(t, async.StatusWorking, future.Status())
}

func TestRunShouldNotHaveValueUntilIsReturnedInChannel(t *testing.T) {
	c := make(chan int)
	future := async.Run(func() int {
		return <-c
	})
	result := future.MustValue()
	require.IsType(t, monad.Error[int]{}, result)
	require.IsType(t, async.FutureNotFinishedError{}, result.TryError())
	c <- 111
	require.Equal(t, 111, future.Value())
	require.Equal(t, async.StatusFinished, future.Status())
}
