// Package stack provides a generic stack (LIFO - Last In, First Out)
// implementation for Go.
//
// The Stack type supports standart operations such as Push, Pop, Peek, and utility
// methods like Clone, Reverse, ToSlice, and Clear.
//
// It is implemented as a wrapper around a Go slice, supporting any comparable type
// T, and offers dynamic resizing as elements are added or removed.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(10)
//	s.Push(20)
//	s.Push(30)
//	fmt.Println(s) // Stack: [10 20 30]
//	top, _ := s.Peek()
//	fmt.Println("Top:", top) // Top: 30
//	value, _ := s.Pop()
//	fmt.Println("Popped:", value) // Popped: 30
//	fmt.Println(s.IsEmpty()) // false
//	s.Clear()
//	fmt.Println(s.IsEmpty()) // true
package stack

import (
	"errors"
	"fmt"
	"slices"
)

// Stack is a generic LIFO (Last In, First Out) data structure.
//
// Stack[T] holds elements of any comparable type T.
//
// Internally, it uses a dynamically growing slice to store elements.
type Stack[T comparable] struct {
	data []T
}

// NewStack creates and returns a new empty Stack for type T.
//
// Returns:
//   - *Stack[T]: A new empty stack for type T.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(42)
//	fmt.Println(s) // Stack: [42]
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

// Push adds a new element to the top of the stack.
//
// Parameters:
//   - data: The element to be added to the stack.
//
// Example:
//
//	s := stack.NewStack[string]()
//	s.Push("hello")
//	s.Push("world")
func (s *Stack[T]) Push(data T) {
	s.data = append(s.data, data)
}

// Pop removes and returns the top element of the stack.
//
// Returns:
//   - value: The top element of the stack.
//   - error: An error if the stack is empty.
//
// If the stack is empty, Pop returns the zero value of T and an error.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(1)
//	value, err := s.Pop()
//	if err == nil {
//	    fmt.Println(value) // 1
//	}
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack empty")
	}
	index := len(s.data) - 1
	value := s.data[index]
	s.data = s.data[:index]
	return value, nil
}

// Peek returns the top element of the stack without removing it.
//
// Returns:
//   - value: The top element of the stack.
//   - error: An error if the stack is empty.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(5)
//	top, err := s.Peek()
//	if err == nil {
//	    fmt.Println(top) // 5
//	}
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack empty")
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty reports whether the stack contains no elements.
//
// Returns:
//   - bool: true if the stack is empty; false otherwise.
//
// Example:
//
//	s := stack.NewStack[int]()
//	fmt.Println(s.IsEmpty()) // true
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements currently in the stack.
//
// Returns:
//   - int: The count of elements in the stack.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(1)
//	s.Push(2)
//	fmt.Println(s.Size()) // 2
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// Clear removes all elements from the stack, resetting it to empty.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(1)
//	s.Clear()
//	fmt.Println(s.IsEmpty()) // true
func (s *Stack[T]) Clear() {
	s.data = nil
}

// String returns a string representation of the stack.
//
// Returns:
//   - string: A string representation of the stack.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(1)
//	s.Push(2)
//	fmt.Println(s.String()) // Stack: [1 2]
func (s *Stack[T]) String() string {
	return fmt.Sprintf("Stack: %v", s.data)
}

// Clone creates and returns a deep copy of the stack.
//
// Returns:
//   - *Stack[T]: A new stack with the same elements.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(1)
//	clone := s.Clone()
//	fmt.Println(clone) // Stack: [1]
func (s *Stack[T]) Clone() *Stack[T] {
	newData := make([]T, len(s.data))
	copy(newData, s.data)
	return &Stack[T]{data: newData}
}

// Reverse reverses the order of elements in the stack.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(1)
//	s.Push(2)
//	s.Push(3)
//	s.Reverse()
//	fmt.Println(s) // Stack: [3 2 1]
func (s *Stack[T]) Reverse() {
	n := s.Size() - 1
	for i, j := 0, n; i < j; i, j = i+1, j-1 {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
}

// ToSlice returns a copy of the stack's elements as a slice.
//
// Returns:
//   - []T: A copy of the stack's internal slice.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(1)
//	s.Push(2)
//	slice := s.ToSlice()
//	fmt.Println(slice) // [1 2]
func (s *Stack[T]) ToSlice() []T {
	result := make([]T, s.Size())
	copy(result, s.data)
	return result
}

// Capacity returns the current capacity of the underlying slice.
//
// Returns:
//   - int: The capacity of the stack's internal slice.
//
// Example:
//
//	s := stack.NewStack[int]()
//	fmt.Println(s.Capacity()) // 0 (initially)
func (s *Stack[T]) Capacity() int {
	return cap(s.data)
}

// Contains reports whether the stack contains the given value.
//
// Parameters:
//   - value: The value to search for.
//
// Returns:
//   - bool: true if the value exists in the stack; false otherwise.
//
// Example:
//
//	s := stack.NewStack[int]()
//	s.Push(10)
//	fmt.Println(s.Contains(10)) // true
//	fmt.Println(s.Contains(5))  // false
func (s *Stack[T]) Contains(value T) bool {
	return slices.Contains(s.data, value)
}
