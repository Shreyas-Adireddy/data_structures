package deque

import (
	"testing"
)

func TestDeque(t *testing.T) {
	// Create a new Deque
	deque := New[int]()

	// Test IsEmpty
	if !deque.IsEmpty() {
		t.Error("Expected deque to be empty")
	}

	// Test AddFront
	deque.AddFront(1)
	deque.AddFront(2)
	deque.AddFront(3)
	if deque.Size() != 3 {
		t.Errorf("Expected deque size to be 3, got %d", deque.Size())
	}

	// Test AddRear
	deque.AddRear(4)
	deque.AddRear(5)
	if deque.Size() != 5 {
		t.Errorf("Expected deque size to be 5, got %d", deque.Size())
	}

	// Test PopFront
	if val, err := deque.PopFront(); err != nil || val != 3 {
		t.Errorf("Expected PopFront to return 3, got %v (error: %v)", val, err)
	}
	if val, err := deque.PopFront(); err != nil || val != 2 {
		t.Errorf("Expected PopFront to return 2, got %v (error: %v)", val, err)
	}

	// Test PopRear
	if val, err := deque.PopRear(); err != nil || val != 5 {
		t.Errorf("Expected PopRear to return 5, got %v (error: %v)", val, err)
	}
	if val, err := deque.PopRear(); err != nil || val != 4 {
		t.Errorf("Expected PopRear to return 4, got %v (error: %v)", val, err)
	}

	// Test PeekFront
	if val, err := deque.PeekFront(); err != nil || val != 1 {
		t.Errorf("Expected PeekFront to return 1, got %v (error: %v)", val, err)
	}

	// Test PopFront on an empty deque
	deque.PopRear()
	if _, err := deque.PopFront(); err == nil || err.Error() != "deque is empty" {
		t.Errorf("Expected PopFront to return error 'deque is empty', got %v", err)
	}

	// Test PopRear on an empty deque
	if _, err := deque.PopRear(); err == nil || err.Error() != "deque is empty" {
		t.Errorf("Expected PopRear to return error 'deque is empty', got %v", err)
	}

	// Test Clear
	deque.Clear()
	if deque.Size() != 0 {
		t.Errorf("Expected deque size to be 0 after Clear, got %d", deque.Size())
	}

	// Test Resize
	deque.Resize(16)
	if len(deque.data) != 16 {
		t.Errorf("Expected deque data length to be 16 after Resize, got %d", len(deque.data))
	}
	deque.Resize(4)
	if len(deque.data) != 4 {
		t.Errorf("Expected deque data length to be 4 after Resize, got %d", len(deque.data))
	}
}
