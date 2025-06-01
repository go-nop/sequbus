package sequbus

import (
	"context"

	"github.com/go-nop/sequbus/runner"
)

// command represents a single command in the command bus
type command[T any] struct {
	handler  runner.Interface[T]
	next     *command[T]
	lastNode *command[T]
}

// newFromRunner creates a new Command from a runner.Interface
func newFromRunner[T any](r runner.Interface[T]) *command[T] {
	return &command[T]{handler: r}
}

// dispatch executes the command and passes the data to the next command in the sequence
// If there is no next command, it simply returns nil
// If an error occurs during execution, it returns the error
func (c *command[T]) dispatch(ctx context.Context, data T) error {
	if err := c.handler.Run(ctx, data); err != nil {
		return err
	}
	if c.next != nil {
		return c.next.dispatch(ctx, data)
	}
	return nil
}
