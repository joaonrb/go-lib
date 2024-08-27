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
	call(err.Err)
	return err
}

func (err Error[T]) WhenOK(call func(T)) Result[T] {
	return err
}

func (err Error[T]) WhenError(call func(error)) Result[T] {
	call(err.Err)
	return err
}

func (err Error[T]) String() string {
	var value T
	return fmt.Sprintf("Error[%T]{Err: \"%v\"}", value, err.Err)
}
