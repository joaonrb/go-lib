package collection

import (
	"context"
	"fmt"
)

type controller[T any] struct {
	context    context.Context
	collection collection[T]
	capacity   uint16
	input      chan chan T
	output     chan chan T
	peek       chan chan T
	length     chan uint16
	flush      chan chan []T
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
		length        uint16
		inputChannel  = ctr.input
		input         = make(chan T)
		outputChannel chan chan T
		output        = make(chan T)
		peekChannel   chan chan T
		peek          = make(chan T)
		flush         = make(chan []T)
	)
	for running := true; running; {
		select {
		case <-ctr.context.Done():
			close(ctr.input)
			close(ctr.output)
			running = false
		case inputChannel <- input:
			value := <-input
			ctr.collection.Push(value)
			length++
			outputChannel = ctr.output
			peekChannel = ctr.peek
			if length >= ctr.capacity {
				inputChannel = nil
			}
		case outputChannel <- output:
			output <- ctr.collection.Pull()
			length--
			inputChannel = ctr.input
			if length == 0 {
				outputChannel = nil
				peekChannel = nil
			}
		case peekChannel <- peek:
			peek <- ctr.collection.Peek()
		case ctr.flush <- flush:
			flush <- ctr.collection.Flush()
			length = 0
		case ctr.length <- length:
		}
	}
}
