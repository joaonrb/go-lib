package op

import "fmt"

type And[T any] struct {
	one, other Evaluator[T]
}

func (and *And[T]) Evaluate(value T) bool {
	return and.one.Evaluate(value) && and.other.Evaluate(value)
}

func (and *And[T]) String() string {
	return fmt.Sprintf("(%s and %s)", and.one, and.other)
}
