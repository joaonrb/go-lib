package collections

import (
	"context"
	"fmt"
	"github.com/joaonrb/go-lib/types"
)

// Queue creates a first in first out Collection with max capacity of 65535.
func Queue[T any](capacity uint16) Collection[T] {
	ctx, cancel := context.WithCancel(context.Background())
	queue := &queue[T]{
		context:  ctx,
		capacity: int(capacity),
		elements: make([]T, capacity),
		input:    make(chan T),
		output:   make(chan T),
		peek:     make(chan T),
		length:   make(chan int),
		flush:    make(chan []T),
		cancel:   cancel,
	}
	queue.start()
	return queue
}

type queue[T any] struct {
	context    context.Context
	capacity   int
	elements   []T
	counter    uint64
	lastInsert int
	lastRead   int
	input      chan T
	output     chan T
	peek       chan T
	length     chan int
	flush      chan []T
	cancel     context.CancelFunc
}

func (collection *queue[T]) Push(value T) types.Option[error] {
	select {
	case <-collection.context.Done():
		return types.Value[error]{This: CollectionClosed()}
	case collection.input <- value:
		return types.Nothing[error]{}
	}
}

func (collection *queue[T]) MustPush(value T) types.Result[bool] {
	select {
	case <-collection.context.Done():
		return types.Error[bool]{Err: CollectionClosed()}
	case collection.input <- value:
		return types.OK[bool]{Value: true}
	default:
		return types.OK[bool]{Value: false}
	}
}

func (collection *queue[T]) Pop() types.Result[T] {
	select {
	case value, open := <-collection.output:
		if !open {
			return types.Error[T]{Err: CollectionClosed()}
		}
		return types.OK[T]{Value: value}
	}
}

func (collection *queue[T]) MustPop() types.Option[T] {
	select {
	case value := <-collection.output:
		return types.Value[T]{This: value}
	default:
		return types.Nothing[T]{}
	}
}

func (collection *queue[T]) Peek() types.Option[T] {
	select {
	case value := <-collection.peek:
		return types.Value[T]{This: value}
	default:
		return types.Nothing[T]{}
	}
}

func (collection *queue[T]) ForEach(ctx context.Context, loop func(index uint64, value T)) {
	var index uint64
	for running := true; running; index++ {
		select {
		case <-ctx.Done():
			running = false
		case value, open := <-collection.output:
			if !open {
				running = false
			} else {
				loop(index, value)
			}
		}
	}
}

func (collection *queue[T]) Flush() types.Result[[]T] {
	select {
	case value, open := <-collection.flush:
		if !open {
			return types.Error[[]T]{Err: CollectionClosed()}
		}
		return types.OK[[]T]{Value: value}
	}
}

func (collection *queue[T]) Length() int {
	return <-collection.length
}

func (collection *queue[T]) IsEmpty() bool {
	return collection.Length() == 0
}

func (collection *queue[T]) Close() {
	collection.cancel()
}

func (collection *queue[T]) IsClosed() bool {
	select {
	case <-collection.context.Done():
		return true
	default:
		return false
	}
}

func (collection *queue[T]) String() string {
	return fmt.Sprintf("Queue{Lenght: %d}", collection.Length())
}

func (collection *queue[T]) start() {
	go collection.controller()
}

func (collection *queue[T]) controller() {
	var (
		length int
		input  = collection.input
		output chan T
		peek   chan T
	)
	for running := true; running; {
		select {
		case <-collection.context.Done():
			close(collection.input)
			close(collection.output)
			running = false
		case value := <-input:
			collection.elements[collection.lastInsert] = value
			length++
			collection.moveInsert()
			output = collection.output
			peek = collection.peek
		case output <- collection.elements[collection.lastRead]:
			length--
			collection.moveRead()
			if length == 0 {
				output = nil
			}
		case peek <- collection.elements[collection.lastRead]:
		case collection.flush <- collection.elementsToSlice():
			length = 0
			collection.lastRead = 0
			collection.lastInsert = 0
		case collection.length <- length:
		}
	}
}

func (collection *queue[T]) moveRead() {
	collection.lastRead = (collection.lastRead + 1) % collection.capacity
}

func (collection *queue[T]) moveInsert() {
	collection.lastInsert = (collection.lastInsert + 1) % collection.capacity
}

func (collection *queue[T]) elementsToSlice() []T {
	switch {
	case collection.lastInsert < collection.lastRead:
		return append(collection.elements[collection.lastRead:], collection.elements[:collection.lastInsert]...)
	case collection.lastInsert > collection.lastRead:
		return collection.elements[collection.lastRead:collection.lastInsert]
	default:
		return nil
	}
}
