package queue

import (
	"testing"
)

func TestBasicEnqueueDequeue(t *testing.T) {
	q := New[int]()
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 5; i++ {
		val, err := q.Dequeue()
		if err != nil || val != i {
			t.Errorf("Dequeue returned unexpected value: %v, expected: %v, error: %v", val, i, err)
		}
	}
}

func TestEnqueueOverflow(t *testing.T) {
	q := New[int]()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 10; i++ {
		val, err := q.Dequeue()
		if err != nil || val != i {
			t.Errorf("Dequeue returned unexpected value: %v, expected: %v, error: %v", val, i, err)
		}
	}
}

func TestDequeueUnderflow(t *testing.T) {
	q := New[int]()
	_, err := q.Dequeue()
	if err == nil || err.Error() != "queue is empty" {
		t.Errorf("Expected 'queue is empty' error, got: %v", err)
	}
}

func TestRepeatedEnqueueDequeue(t *testing.T) {
	q := New[int]()
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 3; i++ {
		_, err := q.Dequeue()
		if err != nil {
			return
		}
	}
	for i := 5; i < 8; i++ {
		q.Enqueue(i)
	}
	for i := 3; i < 8; i++ {
		val, _ := q.Dequeue()
		if val != i {
			t.Errorf("Dequeue returned unexpected value: %v, expected: %v", val, i)
		}
	}
}

func TestEmptyQueueDequeue(t *testing.T) {
	q := New[int]()
	_, err := q.Dequeue()
	if err == nil || err.Error() != "queue is empty" {
		t.Errorf("Expected 'queue is empty' error, got: %v", err)
	}
}

func TestResizeToMinimumCapacity(t *testing.T) {
	q := New[int]()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 5; i++ {
		_, err := q.Dequeue()
		if err != nil {
			t.Errorf("Dequeue returned unexpected error when there should be elements: %v", err)
		}
	}
	for i := 10; i < 15; i++ {
		q.Enqueue(i)
	}
	for i := 5; i < 15; i++ {
		val, _ := q.Dequeue()
		if val != i {
			t.Errorf("Dequeue returned unexpected value: %v, expected: %v", val, i)
		}
	}
}

func TestMixedDataTypes(t *testing.T) {
	qStr := New[string]()
	qStr.Enqueue("hello")
	valStr, _ := qStr.Dequeue()
	if valStr != "hello" {
		t.Errorf("Dequeue returned unexpected value: %v, expected: %v", valStr, "hello")
	}
}

func TestFront(t *testing.T) {
	q := New[int]()
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	front, err := q.Front()
	if err != nil || front != 0 {
		t.Errorf("Front returned unexpected value: %v, expected: %v, error: %v", front, 0, err)
	}
}

func TestBack(t *testing.T) {
	q := New[int]()
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	back, err := q.Back()
	if err != nil || back != 4 {
		t.Errorf("Back returned unexpected value: %v, expected: %v, error: %v", back, 4, err)
	}
}

func TestLenFunction(t *testing.T) {
	q := New[int]()

	// Initially, the queue should be empty
	if length := q.Size(); length != 0 {
		t.Errorf("Expected queue length to be 0, got: %v", length)
	}
	// Enqueue elements into the queue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// After enqueuing, the length should be 3
	if length := q.Size(); length != 3 {
		t.Errorf("Expected queue length to be 3, got: %v", length)
	}

	// Dequeue elements from the queue
	_, err := q.Dequeue()
	if err != nil {
		return
	}
	_, err = q.Dequeue()
	if err != nil {
		return
	}

	// After dequeue, the length should be 1
	if length := q.Size(); length != 1 {
		t.Errorf("Expected queue length to be 1, got: %v", length)
	}
}

func TestMinimumCapacity(t *testing.T) {
	q := New[int]()

	// Enqueue enough elements to trigger resizing
	for i := 0; i < 20; i++ {
		q.Enqueue(i)
	}

	// Dequeue elements until the length of data falls below 4
	for i := 0; i < 16; i++ {
		_, err := q.Dequeue()
		if err != nil {
			return
		}
	}

	// Now, len(data) should be 4
	if capacity := len(q.data); capacity != 8 {
		t.Errorf("Expected queue capacity to be 8 after resizing, got: %v", capacity)
	}
}
