# Stack - Generic Stack Implementation in Go

A fully featured generic **Stack** implementation in Go using slices.

This stack follows the **LIFO** (Last In, First Out) principle and supports core operations along with utility methods for flexibility and convenience.

---

## Table of Contents

- [Stack - Generic Stack Implementation in Go](#stack---generic-stack-implementation-in-go)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Usage](#usage)
  - [Running Tests](#running-tests)
  - [Design Notes](#design-notes)
  - [Example](#example)
  - [Author](#author)
  - [License](#license)
  - [Contact](#contact)

---

## Features

- **Generic**: works with any comparable type (`Stack[T comparable]` in Go 1.18+)

- Core stack operations:

  - `Push(T)` — push an element onto the stack
  - `Pop() (T, error)` — remove and return the top element (returns an error if empty)
  - `Peek() (T, error)` — return (without removing) the top element (error if empty)
  - `IsEmpty() bool` — check if the stack is empty
  - `Size() int` — number of elements in the stack

- Utility methods:

  - `Clear()` — empties the stack
  - `Contains(T) bool` — checks if the stack contains a given element
  - `Clone() *Stack[T]` — creates a deep copy of the stack
  - `Reverse()` — reverses the order of elements in-place
  - `ToSlice() []T` — returns a slice copy of the stack elements
  - `Capacity() int` — returns the capacity of the underlying slice
  - `String() string` — returns a human-readable string representation

- Properly handles empty stack operations by returning Go idiomatic errors (no panics).

- Fully documented with GoDoc/GoComment style comments for `pkg.go.dev`.

---

## Usage

```go
package main

import (
	"fmt"
	"your/module/path/stack" // replace with your actual import path
)

func main() {
	s := stack.NewStack[int]()
	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println(s) // Output: Stack: [10 20 30]

	top, err := s.Peek()
	if err == nil {
		fmt.Println("Top:", top) // Output: Top: 30
	}

	value, err := s.Pop()
	if err == nil {
		fmt.Println("Popped:", value) // Output: Popped: 30
	}

	fmt.Println("Is empty?", s.IsEmpty()) // Output: false

	s.Clear()
	fmt.Println("Is empty after clear?", s.IsEmpty()) // Output: true
}
```

---

## Running Tests

The implementation comes with **comprehensive unit tests** using Go’s `testing` package.

To run all tests:

```bash
go test ./stack -v
```

This will execute tests for:

* Normal stack operations (push, pop, peek).
* Edge cases (empty stack operations).
* Utility methods (`Clone`, `Reverse`, `Contains`, `ToSlice`).

You can also check test coverage:

```bash
go test ./stack -cover
```

---

## Design Notes

- **Internals**: implemented with a Go slice (`[]T`) for dynamic resizing.
- **Generics**: uses Go 1.18+ type parameters (`T comparable`) for flexibility with any comparable type.
- **Error Handling**: all operations that can fail return idiomatic `error` values (no panics).
- **Clone()** creates a deep copy preserving the element order.
- **Reverse()** modifies the stack in-place reversing the element order.
- **String()** implements `fmt.Stringer` for pretty printing.

---

## Example

```go
s := stack.NewStack[string]()
s.Push("Alice")
s.Push("Bob")
s.Push("Charlie")

s.Reverse()

for _, name := range s.ToSlice() {
	fmt.Println(name)
}
// Output:
// Charlie
// Bob
// Alice
```

---

## Author

trigologiaa

---

## License

This project is released under the MIT License. Feel free to use, modify, and distribute.

---

## Contact

For questions or contributions, open an issue or contact the author.