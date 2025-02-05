package atomic

import "fmt"

type Value[T any] struct {
	in  chan T
	out chan T
}

func NewValue[T any](t T) Value[T] {
	value := Value[T]{
		in:  make(chan T),
		out: make(chan T),
	}
	go func() {
		for {
			select {
			case t = <-value.in:
			case value.out <- t:
			}
		}
	}()
	return value
}

func (value Value[T]) Value() T {
	value.ensureIsInitialized()
	return <-value.out
}

func (value Value[T]) Set(v T) {
	value.ensureIsInitialized()
	value.in <- v
}

func (value Value[T]) Swap(n T) (old T) {
	value.ensureIsInitialized()
	old = <-value.out
	value.in <- n
	return
}

func (value Value[T]) String() string {
	value.ensureIsInitialized()
	return fmt.Sprint(value.Value())
}

func (value Value[T]) ensureIsInitialized() {
	if value.in == nil {
		panic(NewValueNotInitializedError())
	}
}
