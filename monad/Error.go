package monad

import (
	"errors"
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

func (err Error[T]) Is(T) bool {
	return false
}

func (err Error[T]) IsIn(...T) bool {
	return false
}

func (err Error[T]) IsError(value error) bool {
	return errors.Is(value, err.Err)
}

func (err Error[T]) IsErrorIn(values ...error) bool {
	result := false
	for i := 0; !result && i < len(values); i++ {
		e := values[i]
		result = result || errors.Is(err.Err, e)
	}
	return result
}

func (err Error[T]) AsError(value any) bool {
	return errors.As(err.Err, &value)
}

func (err Error[T]) AsErrorIn(values ...any) bool {
	result := false
	for i := 0; !result && i < len(values); i++ {
		e := values[i]
		result = result || errors.As(err.Err, &e)
	}
	return result
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
