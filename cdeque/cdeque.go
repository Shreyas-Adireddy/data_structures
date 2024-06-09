package cdeque

import (
	"data_structures/deque"
	"sync"
)

// ConcurrentDeque is a thread-safe double-ended queue.
type ConcurrentDeque[T any] struct {
	dq *deque.Deque[T]
	rw sync.RWMutex
}

// New creates a new ConcurrentDeque.
func New[T any]() *ConcurrentDeque[T] {
	return &ConcurrentDeque[T]{
		dq: deque.New[T](),
	}
}

// IsEmpty checks if the deque is empty.
func (cd *ConcurrentDeque[T]) IsEmpty() bool {
	cd.rw.RLock()
	defer cd.rw.RUnlock()
	return cd.dq.IsEmpty()
}

// Size returns the number of elements in the deque.
func (cd *ConcurrentDeque[T]) Size() int {
	cd.rw.RLock()
	defer cd.rw.RUnlock()
	return cd.dq.Size()
}

// AddFront adds an element to the front of the deque.
func (cd *ConcurrentDeque[T]) AddFront(value T) {
	cd.rw.Lock()
	defer cd.rw.Unlock()
	cd.dq.AddFront(value)
}

// AddRear adds an element to the rear of the deque.
func (cd *ConcurrentDeque[T]) AddRear(value T) {
	cd.rw.Lock()
	defer cd.rw.Unlock()
	cd.dq.AddRear(value)
}

// PopFront removes and returns an element from the front of the deque.
func (cd *ConcurrentDeque[T]) PopFront() (T, error) {
	cd.rw.Lock()
	defer cd.rw.Unlock()
	return cd.dq.PopFront()
}

// PopRear removes and returns an element from the rear of the deque.
func (cd *ConcurrentDeque[T]) PopRear() (T, error) {
	cd.rw.Lock()
	defer cd.rw.Unlock()
	return cd.dq.PopRear()
}

// PeekFront returns the front element without removing it.
func (cd *ConcurrentDeque[T]) PeekFront() (T, error) {
	cd.rw.RLock()
	defer cd.rw.RUnlock()
	return cd.dq.PeekFront()
}

// PeekRear returns the rear element without removing it.
func (cd *ConcurrentDeque[T]) PeekRear() (T, error) {
	cd.rw.RLock()
	defer cd.rw.RUnlock()
	return cd.dq.PeekRear()
}

// Clear removes all elements from the deque.
func (cd *ConcurrentDeque[T]) Clear() {
	cd.rw.Lock()
	defer cd.rw.Unlock()
	cd.dq.Clear()
}

// ToSlice converts the deque to a slice and returns it.
func (cd *ConcurrentDeque[T]) ToSlice() []T {
	cd.rw.RLock()
	defer cd.rw.RUnlock()
	return cd.dq.ToSlice()
}
