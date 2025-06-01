package main

import (
	"context"
	"fmt"

	"github.com/go-nop/sequbus"
	"github.com/go-nop/sequbus/example/entity"
	"github.com/go-nop/sequbus/example/runner"
)

func main() {

	user := &entity.User{
		ID:   "1",
		Name: "Alice",
	}

	bus := sequbus.New[*entity.User]()
	bus.Register(runner.ValidateUser{})
	bus.Register(runner.WelcomeUser{})

	err := bus.Dispatch(context.Background(), user)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
