package queue

import (
	"context"

	"github.com/joaonrb/go-lib/types"
)

func newQueue[T any](manager manager[T], capacity uint16) *Queue[T] {
	ctx, cancel := context.WithCancel(context.Background())
	queue := &Queue[T]{
		context:  ctx,
		manager:  manager,
		capacity: int(capacity),
		elements: make([]T, capacity),
		input:    make(chan T),
		output:   make(chan T),
		peek:     make(chan T),
		length:   make(chan int),
		flush:    make(chan []T),
		cancel:   cancel,
	}
	queue.manager.integrate(queue)
	queue.start()
	return queue
}

type Queue[T any] struct {
	context context.Context
	manager[T]
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

func (queue *Queue[T]) Push(value T) types.Option[error] {
	select {
	case <-queue.context.Done():
		return types.Value[error]{This: CollectionClosed()}
	case queue.input <- value:
		return types.Nothing[error]{}
	}
}

func (queue *Queue[T]) MustPush(value T) types.Result[bool] {
	select {
	case <-queue.context.Done():
		return types.Error[bool]{Err: CollectionClosed()}
	case queue.input <- value:
		return types.OK[bool]{Value: true}
	default:
		return types.OK[bool]{Value: false}
	}
}

func (queue *Queue[T]) Pop() types.Result[T] {
	value, open := <-queue.output
	if !open {
		return types.Error[T]{Err: CollectionClosed()}
	}
	return types.OK[T]{Value: value}
}

func (queue *Queue[T]) MustPop() types.Option[T] {
	select {
	case value := <-queue.output:
		return types.Value[T]{This: value}
	default:
		return types.Nothing[T]{}
	}
}

func (queue *Queue[T]) Peek() types.Option[T] {
	select {
	case value := <-queue.peek:
		return types.Value[T]{This: value}
	default:
		return types.Nothing[T]{}
	}
}

func (queue *Queue[T]) ForEach(ctx context.Context, loop func(index uint64, value T)) {
	var index uint64
	for running := true; running; index++ {
		select {
		case <-ctx.Done():
			running = false
		case value, open := <-queue.output:
			if !open {
				running = false
			} else {
				loop(index, value)
			}
		}
	}
}

func (queue *Queue[T]) Flush() types.Result[[]T] {
	value, open := <-queue.flush
	if !open {
		return types.Error[[]T]{Err: CollectionClosed()}
	}
	return types.OK[[]T]{Value: value}
}

func (queue *Queue[T]) Length() int {
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

func (queue *Queue[T]) start() {
	go queue.controller()
}

func (queue *Queue[T]) controller() {
	var (
		length int
		input  = queue.input
		output chan T
		peek   chan T
	)
	for running := true; running; {
		select {
		case <-queue.context.Done():
			close(queue.input)
			close(queue.output)
			running = false
		case value := <-input:
			queue.manager.insert(value)
			queue.manager.moveInsert()
			length++
			output = queue.output
			peek = queue.peek
			if length >= queue.capacity {
				input = nil
			}
		case output <- queue.manager.read():
			length--
			queue.manager.moveRead()
			input = queue.input
			if length == 0 {
				output = nil
				peek = nil
			}
		case peek <- queue.elements[queue.lastRead]:
		case queue.flush <- queue.manager.flush():
			length = 0
			queue.lastRead = 0
			queue.lastInsert = 0
		case queue.length <- length:
		}
	}
}
