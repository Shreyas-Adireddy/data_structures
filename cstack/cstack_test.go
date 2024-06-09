package cstack

import (
	"sync"
	"testing"
	"time"
)

func TestConcurrentStack(t *testing.T) {
	t.Run("Push and Pop", func(t *testing.T) {
		cs := New[int]()
		wg := sync.WaitGroup{}
		wg.Add(2)

		// Push goroutine
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				cs.Push(i)
			}
		}()

		// Pop goroutine
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				_, err := cs.Pop()
				if err != nil {
					i--
				}
			}
		}()

		wg.Wait()
		if !cs.IsEmpty() {
			t.Errorf("Expected stack to be empty, but it's not")
		}
	})

	t.Run("PopWait", func(t *testing.T) {
		cs := New[int]()
		wg := sync.WaitGroup{}
		wg.Add(2)

		// Push goroutine
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				cs.Push(i)
				time.Sleep(1 * time.Millisecond)
			}
		}()

		// PopWait goroutine
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				_, err := cs.PopWait(10 * time.Millisecond)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		}()

		wg.Wait()
		if !cs.IsEmpty() {
			t.Errorf("Expected stack to be empty, but it's not")
		}
	})

	t.Run("Peek", func(t *testing.T) {
		cs := New[int]()
		cs.Push(1)
		cs.Push(2)
		cs.Push(3)

		top, err := cs.Peek()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if top != 3 {
			t.Errorf("Expected top element to be 3, but got %d", top)
		}

		if cs.Size() != 3 {
			t.Errorf("Expected size to be 3, but got %d", cs.Size())
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		cs := New[int]()
		if !cs.IsEmpty() {
			t.Errorf("Expected stack to be empty, but it's not")
		}

		cs.Push(1)
		if cs.IsEmpty() {
			t.Errorf("Expected stack not to be empty")
		}

		_, _ = cs.Pop()
		if !cs.IsEmpty() {
			t.Errorf("Expected stack to be empty, but it's not")
		}
	})
}

// Tests if stack works as intended
func TestStack(t *testing.T) {
	// Create a new concurrent stack
	s := New[int]()

	// Test Push and Pop
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Test Peek
	top, err := s.Peek()
	if err != nil {
		t.Errorf("Peek error: %v", err)
	}
	if top != 3 {
		t.Errorf("Expected top element to be 3, got %v", top)
	}

	// Test Size
	size := s.Size()
	if size != 3 {
		t.Errorf("Expected size to be 3, got %v", size)
	}

	// Test Pop
	elem, err := s.Pop()
	if err != nil {
		t.Errorf("Pop error: %v", err)
	}
	if elem != 3 {
		t.Errorf("Expected popped element to be 3, got %v", elem)
	}

	// Test IsEmpty after popping
	isEmpty := s.IsEmpty()
	if isEmpty {
		t.Errorf("Expected stack not to be empty after popping, but it is")
	}

	// Test Pop until empty
	for i := 0; i < 2; i++ {
		_, err := s.Pop()
		if err != nil {
			t.Errorf("Pop error: %v", err)
		}
	}

	// Test IsEmpty
	isEmpty = s.IsEmpty()
	if !isEmpty {
		t.Errorf("Expected stack to be empty, but it's not")
	}

	// Test Pop when stack is empty
	_, err = s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack, but got nil")
	}
}
