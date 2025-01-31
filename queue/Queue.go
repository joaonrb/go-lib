package queue

import (
	"context"
	"fmt"

	"github.com/joaonrb/go-lib/monad"
)

// New creates a first in first out New with max capacity of 65535.
func New[T any](capacity uint64) *Queue[T] {
	q := &container[T]{
		capacity: capacity,
		elements: make([]T, capacity),
	}
	queue, ctr := constructor(q, capacity)
	ctr.start()
	return queue
}

func constructor[T any](container *container[T], capacity uint64) (*Queue[T], *controller[T]) {
	ctx, cancel := context.WithCancel(context.Background())
	input := make(chan chan T)
	output := make(chan chan T)
	peek := make(chan chan T)
	length := make(chan uint64)
	flush := make(chan chan []T)
	queue := &Queue[T]{
		context: ctx,
		input:   input,
		output:  output,
		peek:    peek,
		length:  length,
		flush:   flush,
		cancel:  cancel,
	}
	ctr := &controller[T]{
		context:   ctx,
		container: container,
		capacity:  capacity,
		input:     input,
		output:    output,
		peek:      peek,
		length:    length,
		flush:     flush,
	}
	return queue, ctr
}

type Queue[T any] struct {
	context context.Context
	input   chan chan T
	output  chan chan T
	peek    chan chan T
	length  chan uint64
	flush   chan chan []T
	cancel  context.CancelFunc
}

func (queue *Queue[T]) Push(value T) monad.Maybe[error] {
	select {
	case <-queue.context.Done():
		return monad.Some[error]{Value: NewQueueClosedError()}
	case input := <-queue.input:
		input <- value
		return monad.Nothing[error]{}
	}
}

func (queue *Queue[T]) MustPush(value T) monad.Result[bool] {
	select {
	case <-queue.context.Done():
		return monad.Error[bool]{Err: NewQueueClosedError()}
	case input := <-queue.input:
		input <- value
		return monad.OK[bool]{Value: true}
	default:
		return monad.OK[bool]{Value: false}
	}
}

func (queue *Queue[T]) Pull() monad.Result[T] {
	output, open := <-queue.output
	if !open {
		return monad.Error[T]{Err: NewQueueClosedError()}
	}
	return monad.OK[T]{Value: <-output}
}

func (queue *Queue[T]) MustPull() monad.Maybe[T] {
	output, open := <-queue.output
	if open {
		return monad.Some[T]{Value: <-output}
	}
	return monad.Nothing[T]{}
}

func (queue *Queue[T]) Peek() monad.Maybe[T] {
	output, open := <-queue.peek
	if open {
		return monad.Some[T]{Value: <-output}
	}
	return monad.Nothing[T]{}
}

func (queue *Queue[T]) ForEach(ctx context.Context, loop func(index int, value T)) {
	var index int
	for {
		// First select prevent competing goroutines to compete to read output after context is done
		select {
		case <-ctx.Done():
			return
		default:
			select {
			case <-ctx.Done():
				return
			case output, open := <-queue.output:
				if !open {
					return
				} else {
					loop(index, <-output)
					index++
				}
			}
		}
	}
}

func (queue *Queue[T]) ForN(n uint64, loop func(index int, value T)) {
	nn := int(n) - 1
	ctx, cancel := context.WithCancel(context.Background())
	queue.ForEach(ctx, func(index int, value T) {
		loop(index, value)
		if index == nn {
			fmt.Println("cancel", index, value)
			cancel()
			return
		}
	})
}

func (queue *Queue[T]) Flush() monad.Result[[]T] {
	value, open := <-queue.flush
	if !open {
		return monad.Error[[]T]{Err: NewQueueClosedError()}
	}
	return monad.OK[[]T]{Value: <-value}
}

func (queue *Queue[T]) Length() uint64 {
	return <-queue.length
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.Length() == 0
}

func (queue *Queue[T]) Close() {
	queue.cancel()
}

func (queue *Queue[T]) IsClosed() bool {
	select {
	case <-queue.context.Done():
		return true
	default:
		return false
	}
}

func (queue *Queue[T]) String() string {
	var t T
	return fmt.Sprintf("Queue[%T]{Length: %d}", t, queue.Length())
}
