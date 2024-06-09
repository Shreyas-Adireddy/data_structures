package queue

import "fmt"

const maxInt = int(^uint(0) >> 1)

type Queue[T any] struct {
	data  []T
	front int
	rear  int
	size  int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		data:  make([]T, 8),
		front: 0,
		rear:  0,
		size:  0,
	}
}

// Enqueue appends to the rear of the queue.
func (q *Queue[T]) Enqueue(value T) {
	if q.size == maxInt {
		panic("queue size (which is an int) is going to overflow")
	}
	if q.size == len(q.data) {
		newCapacity := q.size
		if q.size < 256 {
			// If the queue is full and small, double the capacity
			newCapacity += q.size
		} else {
			// If the queue is full and large, gradually go from 2x -> 1.25x capacity
			newCapacity += (q.size + 3*256) >> 2
		}
		if newCapacity < 0 {
			// If overflowed, set it to max int
			newCapacity = maxInt
		}
		q.Resize(newCapacity)
	}
	q.data[q.rear] = value
	q.rear = (q.rear + 1) % len(q.data)
	q.size++
}

// Dequeue pops from the front of the queue. If the size becomes less than 1/4 of the capacity, half the capacity.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.size == 0 {
		var null T
		return null, fmt.Errorf("queue is empty")
	}
	value := q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	q.size--
	// If the size becomes less than 1/4 of the capacity, half the capacity
	if q.size > 0 && q.size <= len(q.data)/4 {
		newCapacity := len(q.data) / 2
		if newCapacity < 8 {
			newCapacity = 8
		}
		q.Resize(newCapacity)
	}
	return value, nil
}

func (q *Queue[T]) Front() (T, error) {
	if q.size == 0 {
		var null T
		return null, fmt.Errorf("queue is empty")
	}
	return q.data[q.front], nil
}

func (q *Queue[T]) Back() (T, error) {
	if q.size == 0 {
		var null T
		return null, fmt.Errorf("queue is empty")
	}
	return q.data[(q.rear-1+len(q.data))%len(q.data)], nil
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Clear() {
	q.data = make([]T, 8)
	q.front = 0
	q.rear = 0
	q.size = 0
}

// ToSlice converts the queue to a slice and returns it. It does not make copies of the data within
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, q.size)
	for i := 0; i < q.size; i++ {
		result[i] = q.data[(q.front+i)%len(q.data)]
	}
	return result
}

// Resize adjusts the capacity of the queue. However, Dequeue may half the capacity if queue in not more than 1/4-th full.
func (q *Queue[T]) Resize(newCapacity int) {
	if newCapacity < len(q.data) {
		newCapacity = len(q.data)
	}
	newData := make([]T, newCapacity)
	for i := 0; i < q.size; i++ {
		newData[i] = q.data[(q.front+i)%len(q.data)]
	}
	q.data = newData
	q.front = 0
	q.rear = q.size
}
