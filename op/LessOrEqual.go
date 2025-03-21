package op

import (
	"cmp"
	"fmt"
)

type LessOrEqualEvaluator[T cmp.Ordered] struct {
	value T
}

func LessOrEqual[T cmp.Ordered](value T) Operator[T] {
	evaluator := LessOrEqualEvaluator[T]{value: value}
	return Operator[T]{Evaluator: &evaluator}
}

func (evaluator *LessOrEqualEvaluator[T]) Evaluate(other T) bool {
	return other <= evaluator.value
}

func (evaluator *LessOrEqualEvaluator[T]) String() string {
	return fmt.Sprintf("<= %v", evaluator.value)
}
