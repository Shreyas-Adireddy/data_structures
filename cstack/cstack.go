package cstack

import (
	"data_structures/stack"
	"sync"
)

type ConcurrentStack[T any] struct {
	stack *stack.Stack[T]
	rw    sync.RWMutex // RWMutex for read/write lock
}

// New creates a new concurrent Stack.
func New[T any]() *ConcurrentStack[T] {
	return &ConcurrentStack[T]{
		stack: stack.New[T](),
	}
}

// Push adds an element to the stack.
func (cs *ConcurrentStack[T]) Push(element T) {
	cs.rw.Lock()
	defer cs.rw.Unlock()
	cs.stack.Push(element)
}

// Pop removes and returns the top element of the stack. Blocks if it can't obtain the lock.
func (cs *ConcurrentStack[T]) Pop() (T, error) {
	cs.rw.Lock()
	defer cs.rw.Unlock()
	return cs.stack.Pop()
}

// Peek returns the top element of the stack without removing it.
func (cs *ConcurrentStack[T]) Peek() (T, error) {
	cs.rw.RLock()
	defer cs.rw.RUnlock()
	return cs.stack.Peek()
}

// IsEmpty checks if the stack is empty.
func (cs *ConcurrentStack[T]) IsEmpty() bool {
	cs.rw.RLock()
	defer cs.rw.RUnlock()
	return cs.stack.IsEmpty()
}

// Size returns the number of elements in the stack.
func (cs *ConcurrentStack[T]) Size() int {
	cs.rw.RLock()
	defer cs.rw.RUnlock()
	return cs.stack.Size()
}

// Clear removes all elements from the stack
func (cs *ConcurrentStack[T]) Clear() {
	cs.rw.Lock()
	defer cs.rw.Unlock()
	cs.stack.Clear()
}
