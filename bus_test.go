package sequbus

import (
	"context"
	"errors"
	"testing"
)

// mockRunner implements runner.Interface
type mockRunner[T any] struct {
	name    string
	called  *int
	err     error
	callSeq *[]string
}

func (m *mockRunner[T]) Run(ctx context.Context, data T) error {
	*m.called++
	*m.callSeq = append(*m.callSeq, m.name)
	return m.err
}

func TestCommandBus_RunInOrder(t *testing.T) {
	type MyData struct {
		Value string
	}
	data := &MyData{Value: "test"}

	var counter int
	var sequence []string

	cmd1 := &mockRunner[*MyData]{name: "cmd1", called: &counter, callSeq: &sequence}
	cmd2 := &mockRunner[*MyData]{name: "cmd2", called: &counter, callSeq: &sequence}
	cmd3 := &mockRunner[*MyData]{name: "cmd3", called: &counter, callSeq: &sequence}

	bus := New[*MyData]()
	bus.Register(cmd1)
	bus.Register(cmd2)
	bus.Register(cmd3)

	err := bus.Dispatch(context.Background(), data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if counter != 3 {
		t.Errorf("expected 3 calls, got %d", counter)
	}

	expected := []string{"cmd1", "cmd2", "cmd3"}
	for i, name := range expected {
		if sequence[i] != name {
			t.Errorf("expected %s at position %d, got %s", name, i, sequence[i])
		}
	}
}

func TestCommandBus_ErrorShortCircuits(t *testing.T) {
	type MyData struct{}
	data := &MyData{}

	var counter int
	var sequence []string

	cmd1 := &mockRunner[*MyData]{name: "cmd1", called: &counter, callSeq: &sequence}
	cmd2 := &mockRunner[*MyData]{name: "cmd2", called: &counter, callSeq: &sequence, err: errors.New("fail")}
	cmd3 := &mockRunner[*MyData]{name: "cmd3", called: &counter, callSeq: &sequence}

	bus := New[*MyData]()
	bus.Register(cmd1)
	bus.Register(cmd2)
	bus.Register(cmd3)

	err := bus.Dispatch(context.Background(), data)
	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if counter != 2 {
		t.Errorf("expected 2 calls, got %d", counter)
	}

	expected := []string{"cmd1", "cmd2"}
	for i, name := range expected {
		if sequence[i] != name {
			t.Errorf("expected %s at position %d, got %s", name, i, sequence[i])
		}
	}
}

func TestCommandBus_EmptyDispatch(t *testing.T) {
	type MyData struct{}
	data := &MyData{}

	bus := New[*MyData]()
	err := bus.Dispatch(context.Background(), data)
	if err == nil {
		t.Fatal("expected error when dispatching with no commands")
	}
}
