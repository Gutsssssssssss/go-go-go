package ds

import "errors"

// Generic stack implementation.
type Stack[T any] struct {
	items []T
}

// NewStack returns a new stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack.
// If the stack is empty, it returns an error.
func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Top returns the top item from the stack.
// If the stack is empty, it returns an zero value.
func (s *Stack[T]) Top() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
