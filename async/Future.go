package async

import (
	"fmt"

	"github.com/joaonrb/go-lib/atomic"
	"github.com/joaonrb/go-lib/monad"
)

type Future[T any] struct {
	c      chan T
	e      chan error
	status atomic.Value[Status]
}

func (future Future[T]) Wait() {
	future.checkIfStarted()
	select {
	case value := <-future.c:
		fmt.Println("waited for value", value)
	case value := <-future.e:
		fmt.Println("waited for error", value)
	}
}

func (future Future[T]) Value() T {
	future.checkIfStarted()
	select {
	case value := <-future.c:
		return value
	case value := <-future.e:
		panic(value)
	}
}

func (future Future[T]) MustValue() monad.Result[T] {
	future.checkIfStarted()
	select {
	case value := <-future.c:
		return monad.OK[T]{Value: value}
	case value := <-future.e:
		return monad.Error[T]{Err: value}
	default:
		return monad.Error[T]{Err: NewFutureNotFinishedError()}
	}
}

func (future Future[T]) Channel() <-chan T {
	future.checkIfStarted()
	return future.c
}

func (future Future[T]) ErrorChannel() <-chan error {
	future.checkIfStarted()
	return future.e
}

func (future Future[T]) Status() Status {
	if future.c == nil {
		return StatusNotStarted
	}
	return future.status.Value()
}

func (future Future[T]) String() string {
	var t T
	switch status := future.Status(); status {
	case StatusFinished:
		return fmt.Sprintf("Future[%T]{Status: %s, Value: %v}", t, status, <-future.c)
	case StatusError:
		return fmt.Sprintf("Future[%T]{Status: %s, Error: %v}", t, status, <-future.e)
	default:
		return fmt.Sprintf("Future[%T]{Status: %s}", t, status)
	}
}

func (future Future[T]) checkIfStarted() {
	if future.Status() == StatusNotStarted {
		panic(NewFutureNotStartedError())
	}
}
