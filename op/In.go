package op

import (
	"fmt"
)

type inEquatableEvaluator[T any] struct {
	values []T
}

func (evaluator *inEquatableEvaluator[T]) Evaluate(other T) bool {
	var v any
	for _, v = range evaluator.values {
		if v.(Equatable[T]).Equal(other) {
			return true
		}
	}
	return false
}

func (evaluator *inEquatableEvaluator[T]) String() string {
	return fmt.Sprintf("in %v", evaluator.values)
}

type inDefaultEvaluator[T any] struct {
	values []T
}

func (evaluator *inDefaultEvaluator[T]) Evaluate(other T) bool {
	var (
		v any
		o any = other
	)
	for _, v = range evaluator.values {
		if v == o {
			return true
		}
	}
	return false
}

func (evaluator *inDefaultEvaluator[T]) String() string {
	return fmt.Sprintf("in %v", evaluator.values)
}

func In[T any](values ...T) Operator[T] {
	var value any = values[0]
	switch value.(type) {
	case Equatable[T]:
		return Operator[T]{Evaluator: &inEquatableEvaluator[T]{values: values}}
	default:
		return Operator[T]{Evaluator: &inDefaultEvaluator[T]{values: values}}
	}
}
