package op

import (
	"cmp"
	"fmt"
)

type LessEvaluator[T cmp.Ordered] struct {
	value T
}

func Less[T cmp.Ordered](value T) Operator[T] {
	evaluator := LessEvaluator[T]{value: value}
	return Operator[T]{Evaluator: &evaluator}
}

func (evaluator *LessEvaluator[T]) Evaluate(other T) bool {
	return other < evaluator.value
}

func (evaluator *LessEvaluator[T]) String() string {
	return fmt.Sprintf("< %v", evaluator.value)
}
