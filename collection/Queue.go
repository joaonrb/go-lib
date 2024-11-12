package collection

import "fmt"

// NewQueue creates a first in first out NewQueue with max capacity of 65535.
func NewQueue[T any](capacity uint64) *Queue[T] {
	q := &queue[T]{
		capacity: capacity,
		elements: make([]T, capacity),
	}
	abs, ctr := constructor(q, capacity)
	ctr.start()
	return &Queue[T]{abstract: abs}
}

type Queue[T any] struct {
	*abstract[T]
}

func (collection *Queue[T]) String() string {
	var t T
	return fmt.Sprintf("Queue[%T]{Length: %d}", t, collection.Length())
}

type queue[T any] struct {
	elements []T
	lastPull uint64
	lastPush uint64
	capacity uint64
}

func (collection *queue[T]) Push(value T) {
	collection.elements[collection.lastPush] = value
	collection.lastPush = (collection.lastPush + 1) % collection.capacity
}

func (collection *queue[T]) Pull() T {
	defer func() {
		collection.lastPull = (collection.lastPull + 1) % collection.capacity
	}()
	return collection.elements[collection.lastPull]
}

func (collection *queue[T]) Peek() T {
	return collection.elements[collection.lastPull]
}

func (collection *queue[T]) Flush() []T {
	switch {
	case collection.lastPush < collection.lastPull:
		return append(
			collection.elements[collection.lastPull:],
			collection.elements[:collection.lastPush]...)
	case collection.lastPush > collection.lastPull:
		return collection.elements[collection.lastPull:collection.lastPush]
	default:
		return nil
	}
}

func (collection *queue[T]) String() string {
	var t T
	return fmt.Sprintf("queue[%T]{}", t)
}
