package op

import "cmp"

type Operator[T any] struct {
	Evaluator[T]
}

func (operator Operator[T]) And(other Operator[T]) Operator[T] {
	return func(value T) bool {
		return operator.Evaluate(value) && other.Evaluate(value)
	}
}

func (operator Operator[T]) Or(other Operator[T]) Operator[T] {
	return func(value T) bool {
		return operator(value) || other(value)
	}
}

func Not[T any](operator Operator[T]) Operator[T] {
	return func(value T) bool {
		return !operator(value)
	}
}

func Greater[T cmp.Ordered](other T) Operator[T] {
	return func(value T) bool {
		return value > other
	}
}

func GreaterOrEqual[T cmp.Ordered](other T) Operator[T] {
	return func(value T) bool {
		return value >= other
	}
}

func Less[T cmp.Ordered](other T) Operator[T] {
	return func(value T) bool {
		return value < other
	}
}

func LessOrEqual[T cmp.Ordered](other T) Operator[T] {

	return func(value T) bool {
		return value <= other
	}
}

func In[T any](others ...T) Operator[T] {
	return func(value T) bool {
		var (
			v          any = value
			comparator Operator[T]
		)
		equatable, ok := v.(Equatable[T])
		if ok {
			comparator = func(t T) bool {
				return equatable.Equal(t)
			}
		} else {
			comparator = func(t T) bool {
				return v == toAny(t)
			}
		}
		for _, other := range others {
			if comparator(other) {
				return true
			}
		}
		return false
	}
}

func toAny(value any) any {
	return value
}
