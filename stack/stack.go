package stack

import "errors"

// Stack represents a stack data structure.
type Stack[T any] struct {
	elements []T
}

// New creates a new Stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{elements: []T{}}
}

// Push adds an element to the stack.
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop removes and returns the top element of the stack.
func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}

// Peek returns the top element of the stack without removing it.
func (s *Stack[T]) Peek() (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

// ToSlice converts the stack to a slice and returns it. It does not make copies of the data within
func (s *Stack[T]) ToSlice() []T {
	size := len(s.elements)
	result := make([]T, size)
	for i := 0; i < size; i++ {
		result[i] = s.elements[i]
	}
	return result
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) Clear() {
	s.elements = s.elements[:0]
}

// Resize changes the capacity of the stack to newCapacity. Ideally doesn't make sense to use.
func (s *Stack[T]) Resize(newCapacity int) {
	if newCapacity < len(s.elements) {
		newCapacity = len(s.elements)
	}
	newElements := make([]T, len(s.elements), newCapacity)
	copy(newElements, s.elements)
	s.elements = newElements
}
