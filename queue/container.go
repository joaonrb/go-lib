package queue

import "fmt"

type container[T any] struct {
	elements []T
	lastPull uint64
	lastPush uint64
	capacity uint64
}

func (c *container[T]) Push(value T) {
	c.elements[c.lastPush] = value
	c.lastPush = (c.lastPush + 1) % c.capacity
}

func (c *container[T]) Pull() T {
	defer func() {
		c.lastPull = (c.lastPull + 1) % c.capacity
	}()
	return c.elements[c.lastPull]
}

func (c *container[T]) Peek() T {
	return c.elements[c.lastPull]
}

func (c *container[T]) Flush() []T {
	switch {
	case c.lastPush < c.lastPull:
		return append(
			c.elements[c.lastPull:],
			c.elements[:c.lastPush]...)
	case c.lastPush > c.lastPull:
		return c.elements[c.lastPull:c.lastPush]
	default:
		return nil
	}
}

func (c *container[T]) String() string {
	var t T
	return fmt.Sprintf("container[%T]{}", t)
}
