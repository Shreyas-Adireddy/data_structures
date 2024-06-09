package cqueue

import (
	"data_structures/queue"
	"sync"
)

// ConcurrentQueue is a thread-safe queue.
type ConcurrentQueue[T any] struct {
	q  *queue.Queue[T]
	rw sync.RWMutex
}

// New creates a new ConcurrentQueue.
func New[T any]() *ConcurrentQueue[T] {
	return &ConcurrentQueue[T]{
		q: queue.New[T](),
	}
}

// Enqueue adds an element to the rear of the queue.
func (cq *ConcurrentQueue[T]) Enqueue(value T) {
	cq.rw.Lock()
	defer cq.rw.Unlock()
	cq.q.Enqueue(value)
}

// Dequeue removes and returns an element from the front of the queue.
func (cq *ConcurrentQueue[T]) Dequeue() (T, error) {
	cq.rw.Lock()
	defer cq.rw.Unlock()
	return cq.q.Dequeue()
}

// Front returns the front element of the queue without removing it.
func (cq *ConcurrentQueue[T]) Front() (T, error) {
	cq.rw.RLock()
	defer cq.rw.RUnlock()
	return cq.q.Front()
}

// Back returns the rear element of the queue without removing it.
func (cq *ConcurrentQueue[T]) Back() (T, error) {
	cq.rw.RLock()
	defer cq.rw.RUnlock()
	return cq.q.Back()
}

// Size returns the number of elements in the queue.
func (cq *ConcurrentQueue[T]) Size() int {
	cq.rw.RLock()
	defer cq.rw.RUnlock()
	return cq.q.Size()
}

// IsEmpty checks if the queue is empty.
func (cq *ConcurrentQueue[T]) IsEmpty() bool {
	cq.rw.RLock()
	defer cq.rw.RUnlock()
	return cq.q.IsEmpty()
}

// Clear removes all elements from the queue.
func (cq *ConcurrentQueue[T]) Clear() {
	cq.rw.Lock()
	defer cq.rw.Unlock()
	cq.q.Clear()
}

// ToSlice converts the queue to a slice and returns it.
func (cq *ConcurrentQueue[T]) ToSlice() []T {
	cq.rw.RLock()
	defer cq.rw.RUnlock()
	return cq.q.ToSlice()
}
