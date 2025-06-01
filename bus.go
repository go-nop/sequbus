package sequbus

import (
	"context"
	"fmt"

	"github.com/go-nop/sequbus/runner"
)

// CommandBus is a sequential command bus that allows registering commands
type CommandBus[T any] struct {
	head *command[T]
}

// New creates a new instance of CommandBus
func New[T any]() *CommandBus[T] {
	return &CommandBus[T]{}
}

// Register adds a new command to the bus
func (cb *CommandBus[T]) Register(r runner.Interface[T]) {
	newNode := newFromRunner(r)

	if cb.head == nil {
		cb.head = newNode
		cb.head.lastNode = newNode
		return
	}

	cb.head.lastNode.next = newNode
	cb.head.lastNode = newNode
}

// Dispatch executes the registered commands in sequence
func (cb *CommandBus[T]) Dispatch(ctx context.Context, data T) error {
	if cb.head == nil {
		return fmt.Errorf("no command registered")
	}
	return cb.head.dispatch(ctx, data)
}
