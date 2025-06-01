package runner

import (
	"context"
	"errors"

	"github.com/go-nop/sequbus/example/entity"
)

type ValidateUser struct{}

// Run validates the user entity
func (v ValidateUser) Run(ctx context.Context, user *entity.User) error {
	if user.ID == "" {
		return errors.New("user ID is required")
	}
	return nil
}
