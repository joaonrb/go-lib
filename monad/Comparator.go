package monad

import "cmp"

type Comparator[T any] func(T) bool

type Equatable[T any] interface {
	Equal(other T) bool
}

func Equal[T any](other T) Comparator[T] {
	return func(value T) bool {
		var v any = value
		equatable, ok := v.(Equatable[T])
		return (ok && equatable.Equal(other)) || v == toAny(other)
	}
}

func Different[T any](other T) Comparator[T] {
	equal := Equal(other)
	return func(value T) bool {
		return !equal(value)
	}
}

func Greater[T cmp.Ordered](other T) Comparator[T] {
	return func(value T) bool {
		return value > other
	}
}

func GreaterOrEqual[T cmp.Ordered](other T) Comparator[T] {
	return func(value T) bool {
		return value >= other
	}
}

func Less[T cmp.Ordered](other T) Comparator[T] {
	return func(value T) bool {
		return value < other
	}
}

func LessOrEqual[T cmp.Ordered](other T) Comparator[T] {
	return func(value T) bool {
		return value <= other
	}
}

func In[T any](others ...T) Comparator[T] {
	return func(value T) bool {
		var v, other any = value, nil
		for _, other = range others {
			if v == other {
				return true
			}
		}
		return false
	}
}

func NotIn[T any](others ...T) Comparator[T] {
	in := In(others...)
	return func(value T) bool {
		return !in(value)
	}
}

func toAny(value any) any {
	return value
}
