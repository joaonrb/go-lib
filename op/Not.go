package op

import "fmt"

type NotEvaluator[T any] struct {
	evaluator Evaluator[T]
}

func Not[T any](operator Operator[T]) Operator[T] {
	evaluator := NotEvaluator[T]{evaluator: operator.Evaluator}
	return Operator[T]{Evaluator: &evaluator}
}

func (evaluator *NotEvaluator[T]) Evaluate(other T) bool {
	return !evaluator.evaluator.Evaluate(other)
}

func (evaluator *NotEvaluator[T]) String() string {
	return fmt.Sprintf("not %s", evaluator.evaluator)
}
