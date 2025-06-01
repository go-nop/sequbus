package runner

import "context"

type Interface[T any] interface {
	Run(ctx context.Context, data T) error
}
