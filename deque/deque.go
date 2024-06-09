package deque

import (
	"errors"
)

const maxInt = int(^uint(0) >> 1)

// Deque represents a double-ended queue implemented with a circular array.
type Deque[T any] struct {
	data  []T
	front int
	rear  int
	size  int
}

// New creates a new Deque with the given initial capacity.
func New[T any]() *Deque[T] {
	return &Deque[T]{
		data:  make([]T, 8),
		front: 0,
		rear:  0,
		size:  0,
	}
}

// IsEmpty checks if the deque is empty.
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// Size returns the number of elements in the deque.
func (d *Deque[T]) Size() int {
	return d.size
}

// AddFront adds an element to the front of the deque.
func (d *Deque[T]) AddFront(value T) {
	if d.size == maxInt {
		panic("deque size (which is an int) is going to overflow")
	}
	if d.size == len(d.data) {
		newCapacity := d.size
		if d.size < 256 {
			// If the queue is full and small, double the capacity
			newCapacity += d.size
		} else {
			// If the queue is full and large, gradually go from 2x -> 1.25x capacity
			newCapacity += (d.size + 3*256) >> 2
		}
		if newCapacity < 0 {
			// If overflowed, set it to max int
			newCapacity = maxInt
		}
		d.Resize(newCapacity)
	}
	d.front = (d.front - 1 + len(d.data)) % len(d.data)
	d.data[d.front] = value
	d.size++
}

// AddRear adds an element to the rear of the deque.
func (d *Deque[T]) AddRear(value T) {
	if d.size == maxInt {
		panic("deque size (which is an int) is going to overflow")
	}
	if d.size == len(d.data) {
		newCapacity := d.size
		if d.size < 256 {
			// If the queue is full and small, double the capacity
			newCapacity += d.size
		} else {
			// If the queue is full and large, gradually go from 2x -> 1.25x capacity
			newCapacity += (d.size + 3*256) >> 2
		}
		if newCapacity < 0 {
			// If overflowed, set it to max int
			newCapacity = maxInt
		}
		d.Resize(newCapacity)
	}
	d.data[d.rear] = value
	d.rear = (d.rear + 1) % len(d.data)
	d.size++
}

// PopFront removes and returns an element from the front of the deque.
func (d *Deque[T]) PopFront() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, errors.New("deque is empty")
	}
	value := d.data[d.front]
	d.front = (d.front + 1) % len(d.data)
	d.size--
	if d.size > 0 && d.size <= len(d.data)/4 {
		newCapacity := len(d.data) / 2
		if newCapacity < 8 {
			newCapacity = 8
		}
		d.Resize(newCapacity)
	}
	return value, nil
}

// PopRear removes and returns an element from the rear of the deque.
func (d *Deque[T]) PopRear() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, errors.New("deque is empty")
	}
	d.rear = (d.rear - 1 + len(d.data)) % len(d.data)
	value := d.data[d.rear]
	d.size--
	if d.size > 0 && d.size <= len(d.data)/4 {
		newCapacity := len(d.data) / 2
		if newCapacity < 8 {
			newCapacity = 8
		}
		d.Resize(newCapacity)
	}
	return value, nil
}

// PeekFront returns the front element without removing it.
func (d *Deque[T]) PeekFront() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, errors.New("deque is empty")
	}
	return d.data[d.front], nil
}

// PeekRear returns the rear element without removing it.
func (d *Deque[T]) PeekRear() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, errors.New("deque is empty")
	}
	rearIndex := (d.rear - 1 + len(d.data)) % len(d.data)
	return d.data[rearIndex], nil
}

// Clear removes all elements from the deque.
func (d *Deque[T]) Clear() {
	d.data = make([]T, 8)
	d.front = 0
	d.rear = 0
	d.size = 0
}

// ToSlice converts the queue to a slice and returns it. It does not make copies of the data within
func (d *Deque[T]) ToSlice() []T {
	result := make([]T, d.size)
	for i := 0; i < d.size; i++ {
		result[i] = d.data[(d.front+i)%len(d.data)]
	}
	return result
}

// Resize resizes the underlying array to the new capacity.
func (d *Deque[T]) Resize(newCapacity int) {
	if newCapacity < 1 {
		newCapacity = 1
	}
	newData := make([]T, newCapacity)
	for i := 0; i < d.size; i++ {
		newData[i] = d.data[(d.front+i)%len(d.data)]
	}
	d.data = newData
	d.front = 0
	d.rear = d.size
}
