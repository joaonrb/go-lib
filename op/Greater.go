package op

import (
	"cmp"
	"fmt"
)

type GreaterEvaluator[T cmp.Ordered] struct {
	value T
}

func Greater[T cmp.Ordered](value T) Operator[T] {
	evaluator := GreaterEvaluator[T]{value: value}
	return Operator[T]{Evaluator: &evaluator}
}

func (evaluator *GreaterEvaluator[T]) Evaluate(other T) bool {
	return other > evaluator.value
}

func (evaluator *GreaterEvaluator[T]) String() string {
	return fmt.Sprintf("> %v", evaluator.value)
}
