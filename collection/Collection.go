package collection

import (
	"context"

	"github.com/joaonrb/go-lib/types"
)

type Collection[T any] interface {
	// Push sends an element to que abstract, blocking while waiting for space if the abstract has no vacancy. Returns a
	// Value with error if the abstract is closed.
	Push(T) types.Option[error]

	// MustPush sends an element to que abstract and returns true. Returns an Ok with false if the abstract has no vacancy or
	// an Error if the abstract is closed.
	MustPush(T) types.Result[bool]

	// Pull remove and returns an OK with the next element in the abstract or Error if the abstract is closed.
	Pull() types.Result[T]

	// MustPull remove and returns a Value with the next element in the abstract or Nothing.
	MustPull() types.Option[T]

	// Peek returns a Value with the next element in the abstract without removing it or Nothing.
	Peek() types.Option[T]

	// ForEach loops over the next unread elements in the abstract.
	ForEach(ctx context.Context, loop func(index int, value T))

	// ForN loops over the next N unread elements in the abstract.
	ForN(n uint16, loop func(index int, value T))

	// Flush removes and returns OK with all elements from the abstract. Returns an Error if the abstract is closed.
	Flush() types.Result[[]T]

	// Length returns the number of elements in the abstract.
	Length() int

	// IsEmpty returns true if the abstract has no elements.
	IsEmpty() bool

	// Close closes the abstract. After the abstract is closed it will not accept new elements. Old elements can be retrieved.
	Close()

	// IsClosed returns true if the abstract is closed.
	IsClosed() bool
}

type collection[T any] interface {
	Push(value T)
	Pull() T
	Peek() T
	Flush() []T
}

func constructor[T any](collection collection[T], capacity uint16) (*abstract[T], *controller[T]) {
	ctx, cancel := context.WithCancel(context.Background())
	input := make(chan chan T)
	output := make(chan chan T)
	peek := make(chan chan T)
	length := make(chan uint16)
	flush := make(chan chan []T)
	abs := &abstract[T]{
		context: ctx,
		input:   input,
		output:  output,
		peek:    peek,
		length:  length,
		flush:   flush,
		cancel:  cancel,
	}
	ctr := &controller[T]{
		context:    ctx,
		collection: collection,
		capacity:   capacity,
		input:      input,
		output:     output,
		peek:       peek,
		length:     length,
		flush:      flush,
	}
	return abs, ctr
}
