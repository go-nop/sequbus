package sequbus

import (
	"context"
	"fmt"
	"testing"
)

// MockRunner is a mock implementation of runner.Interface
type MockRunner[T any] struct {
	shouldFail bool
}

func (m *MockRunner[T]) Run(ctx context.Context, data T) error {
	if m.shouldFail {
		return fmt.Errorf("mock error")
	}
	return nil
}

func TestCommand_Dispatch(t *testing.T) {
	type testData struct {
		value string
	}

	t.Run("dispatch single command successfully", func(t *testing.T) {
		mockRunner := &MockRunner[testData]{shouldFail: false}
		cmd := newFromRunner(mockRunner)

		err := cmd.dispatch(context.Background(), testData{value: "test"})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("dispatch single command with error", func(t *testing.T) {
		mockRunner := &MockRunner[testData]{shouldFail: true}
		cmd := newFromRunner(mockRunner)

		err := cmd.dispatch(context.Background(), testData{value: "test"})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})

	t.Run("dispatch multiple commands successfully", func(t *testing.T) {
		mockRunner1 := &MockRunner[testData]{shouldFail: false}
		mockRunner2 := &MockRunner[testData]{shouldFail: false}

		cmd1 := newFromRunner(mockRunner1)
		cmd2 := newFromRunner(mockRunner2)
		cmd1.next = cmd2

		err := cmd1.dispatch(context.Background(), testData{value: "test"})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	t.Run("dispatch multiple commands with error in second command", func(t *testing.T) {
		mockRunner1 := &MockRunner[testData]{shouldFail: false}
		mockRunner2 := &MockRunner[testData]{shouldFail: true}

		cmd1 := newFromRunner(mockRunner1)
		cmd2 := newFromRunner(mockRunner2)
		cmd1.next = cmd2

		err := cmd1.dispatch(context.Background(), testData{value: "test"})
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}
