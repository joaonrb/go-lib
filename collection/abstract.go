package collection

import (
	"context"
	"fmt"

	"github.com/joaonrb/go-lib/types"
)

type abstract[T any] struct {
	context context.Context
	input   chan chan T
	output  chan chan T
	peek    chan chan T
	length  chan uint64
	flush   chan chan []T
	cancel  context.CancelFunc
}

func (abs *abstract[T]) Push(value T) types.Option[error] {
	select {
	case <-abs.context.Done():
		return types.Value[error]{This: CollectionClosed()}
	case input := <-abs.input:
		input <- value
		return types.Nothing[error]{}
	}
}

func (abs *abstract[T]) MustPush(value T) types.Result[bool] {
	select {
	case <-abs.context.Done():
		return types.Error[bool]{Err: CollectionClosed()}
	case input := <-abs.input:
		input <- value
		return types.OK[bool]{Value: true}
	default:
		return types.OK[bool]{Value: false}
	}
}

func (abs *abstract[T]) Pull() types.Result[T] {
	output, open := <-abs.output
	if !open {
		return types.Error[T]{Err: CollectionClosed()}
	}
	return types.OK[T]{Value: <-output}
}

func (abs *abstract[T]) MustPull() types.Option[T] {
	output, open := <-abs.output
	if open {
		return types.Value[T]{This: <-output}
	}
	return types.Nothing[T]{}
}

func (abs *abstract[T]) Peek() types.Option[T] {
	output, open := <-abs.peek
	if open {
		return types.Value[T]{This: <-output}
	}
	return types.Nothing[T]{}
}

func (abs *abstract[T]) ForEach(ctx context.Context, loop func(index int, value T)) {
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
			case output, open := <-abs.output:
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

func (abs *abstract[T]) ForN(n uint64, loop func(index int, value T)) {
	nn := int(n) - 1
	ctx, cancel := context.WithCancel(context.Background())
	abs.ForEach(ctx, func(index int, value T) {
		loop(index, value)
		if index == nn {
			fmt.Println("cancel", index, value)
			cancel()
			return
		}
	})
}

func (abs *abstract[T]) Flush() types.Result[[]T] {
	value, open := <-abs.flush
	if !open {
		return types.Error[[]T]{Err: CollectionClosed()}
	}
	return types.OK[[]T]{Value: <-value}
}

func (abs *abstract[T]) Length() uint64 {
	return <-abs.length
}

func (abs *abstract[T]) IsEmpty() bool {
	return abs.Length() == 0
}

func (abs *abstract[T]) Close() {
	abs.cancel()
}

func (abs *abstract[T]) IsClosed() bool {
	select {
	case <-abs.context.Done():
		return true
	default:
		return false
	}
}

func (abs *abstract[T]) String() string {
	var t T
	return fmt.Sprintf("abstract[%T]{}", t)
}
