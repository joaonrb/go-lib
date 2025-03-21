package op

import "fmt"

type Or[T any] struct {
	one, other Evaluator[T]
}

func (or *Or[T]) Evaluate(value T) bool {
	return or.one.Evaluate(value) || or.other.Evaluate(value)
}

func (or *Or[T]) String() string {
	return fmt.Sprintf("(%s or %s)", or.one, or.other)
}
