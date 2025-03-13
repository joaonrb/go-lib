package op

import "fmt"

type EqualOperator[T any] struct {
	value    any
	evaluate func(any) bool
}

func Equal[T any](value T) Operator[T] {
	operator := EqualOperator[T]{value: value}
	equatable, ok := operator.value.(Equatable[T])
	if ok {
		operator.evaluate = equatable.Equal
	} else {
		operator.evaluate = operator.defaultEvaluate
	}
	return &operator
}

func (operation *EqualOperator[T]) Evaluate(other T) bool {
	return operation.evaluate(other)
}

func (operation *EqualOperator[T]) defaultEvaluate(other any) bool {
	return operation.value == other
}

func (operation *EqualOperator[T]) String() string {
	return fmt.Sprint()
}
