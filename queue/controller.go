package queue

import (
	"context"
	"fmt"
)

type controller[T any] struct {
	context   context.Context
	container *container[T]
	capacity  uint64
	input     chan chan T
	output    chan chan T
	peek      chan chan T
	length    chan uint64
	flush     chan chan []T
}

func (ctr *controller[T]) String() string {
	var t T
	return fmt.Sprintf("controller[%T]{}", t)
}

func (ctr *controller[T]) start() {
	go ctr.run()
}

func (ctr *controller[T]) run() {
	var (
		length        uint64
		inputChannel  = ctr.input
		input         = make(chan T)
		outputChannel chan chan T
		output        = make(chan T)
		peekChannel   chan chan T
		peek          = make(chan T)
		flush         = make(chan []T)
		done          = ctr.context.Done()
	)
	defer func() {
		close(ctr.input)
		close(ctr.output)
		close(ctr.peek)
		close(ctr.length)
		close(ctr.flush)
	}()
	for running := true; running || length > 0; {
		select {
		case inputChannel <- input:
			value := <-input
			ctr.container.Push(value)
			length++
			outputChannel = ctr.output
			peekChannel = ctr.peek
			if length >= ctr.capacity {
				inputChannel = nil
			}
		case outputChannel <- output:
			output <- ctr.container.Pull()
			length--
			inputChannel = ctr.input
			if length == 0 {
				outputChannel = nil
				peekChannel = nil
			}
		case peekChannel <- peek:
			peek <- ctr.container.Peek()
		case ctr.flush <- flush:
			flush <- ctr.container.Flush()
			length = 0
		case ctr.length <- length:
		case <-done:
			running = false
			done = nil
		}
	}
}
