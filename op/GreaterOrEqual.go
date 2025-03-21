package op

import (
	"cmp"
	"fmt"
)

type GreaterOrEqualEvaluator[T cmp.Ordered] struct {
	value T
}

func GreaterOrEqual[T cmp.Ordered](value T) Operator[T] {
	evaluator := GreaterOrEqualEvaluator[T]{value: value}
	return Operator[T]{Evaluator: &evaluator}
}

func (evaluator *GreaterOrEqualEvaluator[T]) Evaluate(other T) bool {
	return other >= evaluator.value
}

func (evaluator *GreaterOrEqualEvaluator[T]) String() string {
	return fmt.Sprintf(">= %v", evaluator.value)
}
