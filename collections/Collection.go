package collections

import (
	"context"
	"github.com/joaonrb/go-lib/types"
)

type Collection[T any] interface {
	// Push sends an element to que queue, blocking while waiting for space if the queue has no vacancy. Returns a
	// Value with error if the queue is closed.
	Push(T) types.Option[error]

	// MustPush sends an element to que queue and returns true. Returns an Ok with false if the queue has no vacancy or
	// an Error if the queue is closed.
	MustPush(T) types.Result[bool]

	// Pop remove and returns an OK with the next element in the queue or Error if the queue is closed.
	Pop() types.Result[T]

	// MustPop remove and returns a Value with the next element in the queue or Nothing.
	MustPop() types.Option[T]

	// Peek returns a Value with the next element in the queue without removing it or Nothing.
	Peek() types.Option[T]

	// ForEach loops over the unread elements in the queue.
	ForEach(ctx context.Context, loop func(index uint64, value T))

	// Flush removes and returns OK with all elements from the queue. Returns an Error if the queue is closed.
	Flush() types.Result[[]T]

	// Length returns the number of elements in the queue.
	Length() int

	// IsEmpty returns true if the queue has no elements.
	IsEmpty() bool

	// Close closes the queue. After the queue is closed it will not accept new elements. Old elements can be retrieved.
	Close()

	// IsClosed returns true if the queue is closed.
	IsClosed() bool
}
