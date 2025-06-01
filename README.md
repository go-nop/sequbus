# sequbus
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/go-nop/sequbus/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/go-nop/strix)](https://github.com/go-nop/sequbus/releases/latest)
![GolangVersion](https://img.shields.io/github/go-mod/go-version/go-nop/sequbus)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-nop/sequbus.svg)](https://pkg.go.dev/github.com/go-nop/sequbus)

**`sequbus`** is a generic and lightweight sequential command bus for Go. It allows you to register and execute a series of commands in a strict order. Inspired by middleware and pipeline patterns.

## ✨ Features

- ✅ Generic: works with any data type
- ✅ Sequential: commands run one after another
- ✅ Chainable: register multiple commands easily
- ✅ Early exit on error
- ✅ Simple interface

## 📦 Installation

```bash
go get github.com/go-nop/sequbus
```

## 🚀 Usage

### 1. Define Your Data and Command(s)

```go
type User struct {
	ID   string
	Name string
}

type ValidateUser struct{}

func (v ValidateUser) Run(ctx context.Context, user *User) error {
	if user.ID == "" {
		return errors.New("missing ID")
	}
	return nil
}
```

### 2. Register Commands to the Bus

```go
bus := sequbus.New[*User]()
bus.Register(ValidateUser{})
// bus.Register(SendWelcomeEmail{}) // More commands

user := &User{ID: "123", Name: "Alice"}
err := bus.Dispatch(context.Background(), user)
if err != nil {
	log.Fatal(err)
}
```

## 🧪 Testing

Run all tests:

```bash
go test ./...
```

## 📂 Project Structure

```
sequbus/
├── bus.go                 // CommandBus implementation
├── bus_test.go         // Unit tests
├── node.go              // Command node implementation
├── node_test.go      // Unit tests
├── runner/
    └── interface.go   // runner.Interface definition
└── example/
	└── main.go     // Example usage

```

## 📘 Interface

```go
type Interface[T any] interface {
	Run(ctx context.Context, data T) error
}
```

Implement this interface to create custom commands.

---

## 🛠 Example Use Cases

- User registration pipeline
- Data processing workflows
- Approval chains
- Middleware-like systems in CLI or microservices

## 📄 License

MIT

---

Made with ❤️ by [@go-nop](https://github.com/go-nop)
