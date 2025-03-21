package op

import "fmt"

type EqualEvaluator[T any] struct {
	value    any
	evaluate func(any) bool
}

func Equal[T any](value T) Operator[T] {
	evaluator := EqualEvaluator[T]{value: value}
	equatable, ok := evaluator.value.(Equatable[T])
	if ok {
		evaluator.evaluate = equatable.Equal
	} else {
		evaluator.evaluate = evaluator.defaultEvaluate
	}
	return Operator[T]{Evaluator: &evaluator}
}

func (evaluator *EqualEvaluator[T]) Evaluate(other T) bool {
	return evaluator.evaluate(other)
}

func (evaluator *EqualEvaluator[T]) defaultEvaluate(other any) bool {
	return evaluator.value == other
}

func (evaluator *EqualEvaluator[T]) String() string {
	return fmt.Sprintf("X == %s", evaluator.value)
}
