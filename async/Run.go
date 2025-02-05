package async

import (
	"fmt"
	"github.com/joaonrb/go-lib/atomic"
)

func Run[T any](f func() T) Future[T] {
	var promise Future[T]
	run(&promise, f)
	return promise
}

func run[T any](promise *Future[T], f func() T) {
	status := promise.Status()
	if status == StatusFinished || status == StatusError {
		return
	}
	promise.c = make(chan T)
	promise.e = make(chan error)
	promise.status = atomic.NewValue(StatusWorking)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				var (
					e  error
					ok bool
				)
				if e, ok = err.(error); !ok {
					e = fmt.Errorf("%v", err)
				}
				promise.status.Set(StatusError)
				for {
					promise.e <- e
				}
			}
		}()
		value := f()
		promise.status.Set(StatusFinished)
		for {
			promise.c <- value
		}
	}()
}
