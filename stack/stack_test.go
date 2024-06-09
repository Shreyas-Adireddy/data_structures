package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New[int]()

	if !s.IsEmpty() {
		t.Error("Expected stack to be empty")
	}

	s.Push(1)
	s.Push(2)

	if s.Size() != 2 {
		t.Errorf("Expected stack size to be 2, got %d", s.Size())
	}

	top, err := s.Peek()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if top != 2 {
		t.Errorf("Expected top element to be 2, got %v", top)
	}

	popped, err := s.Pop()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if popped != 2 {
		t.Errorf("Expected popped element to be 2, got %v", popped)
	}

	if s.Size() != 1 {
		t.Errorf("Expected stack size to be 1, got %d", s.Size())
	}
}
