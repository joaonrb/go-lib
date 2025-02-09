package async

import (
	"reflect"
)

func race[T any](promises []Future[T], f func(int, Future[T])) {
	cases := make([]reflect.SelectCase, len(promises))
	for i, promise := range promises {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(waitForPromise(promise)),
		}
	}
	for i := 0; i < len(promises); {
		chosen, value, ok := reflect.Select(cases)
		if ok {
			f(chosen, value.Interface().(Future[T]))
			i++
		}
	}
}

func waitForPromise[T any](promise Future[T]) chan Future[T] {
	ch := make(chan Future[T])
	go func() {
		defer close(ch)
		promise.Wait()
		ch <- promise
	}()
	return ch
}
