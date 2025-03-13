package types

import (
	"fmt"
)

var _ Result[any] = Error[any]{}

type Error[T any] struct {
	Err error
}

func (err Error[T]) result() {}

func (err Error[T]) Then(func(T) Result[T]) Result[T] {
	return err
}

func (err Error[T]) Error(call func(error) Result[T]) Result[T] {
	return call(err.Err)
}

func (err Error[T]) WhenOK(call func(T)) Result[T] {
	return err
}

func (err Error[T]) WhenError(call func(error)) Result[T] {
	call(err.Err)
	return err
}

func (err Error[T]) Or(value T) Result[T] {
	return OK[T]{Value: value}
}

func (err Error[T]) TryValue() T {
	panic(NewResultIsErrorError(err.Err))
}

func (err Error[T]) TryError() error {
	return err.Err
}

func (err Error[T]) String() string {
	var value T
	return fmt.Sprintf("Error[%T]{Err: \"%v\"}", value, err.Err)
}
