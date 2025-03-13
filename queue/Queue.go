package queue

import "fmt"

// New creates a first in first out Queue with max capacity of 65535.
func New[T any](capacity uint16) *Queue[T] {
	return newQueue[T](&queueManager[T]{}, capacity)
}

type queueManager[T any] struct {
	queue *Queue[T]
}

func (m *queueManager[T]) read() T {
	return m.queue.elements[m.queue.lastRead]
}

func (m *queueManager[T]) moveRead() {
	m.queue.lastRead = (m.queue.lastRead + 1) % m.queue.capacity
}

func (m *queueManager[T]) insert(value T) {
	m.queue.elements[m.queue.lastInsert] = value
}

func (m *queueManager[T]) moveInsert() {
	m.queue.lastInsert = (m.queue.lastInsert + 1) % m.queue.capacity
}

func (m *queueManager[T]) flush() []T {
	switch {
	case m.queue.lastInsert < m.queue.lastRead:
		return append(m.queue.elements[m.queue.lastRead:], m.queue.elements[:m.queue.lastInsert]...)
	case m.queue.lastInsert > m.queue.lastRead:
		return m.queue.elements[m.queue.lastRead:m.queue.lastInsert]
	default:
		return nil
	}
}

func (m *queueManager[T]) integrate(queue *Queue[T]) {
	m.queue = queue
}

func (m *queueManager[T]) String() string {
	var t T
	return fmt.Sprintf("New[%T]{Lenght: %d}", t, m.queue.Length())
}
