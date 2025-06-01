package runner

import (
	"context"
	"fmt"

	"github.com/go-nop/sequbus/example/entity"
)

type WelcomeUser struct{}

// Run welcomes the user by printing a message
func (w WelcomeUser) Run(ctx context.Context, user *entity.User) error {
	fmt.Printf("Welcome, %s!\n", user.Name)
	return nil
}
