package stack

import "testing"

func TestPushAndPop(t *testing.T) {
	s := NewStack[int]()
	if !s.IsEmpty() {
		t.Error("expected stack to be empty initially")
	}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	if s.Size() != 3 {
		t.Errorf("expected size 3, got %d", s.Size())
	}
	if s.IsEmpty() {
		t.Error("expected stack to be non-empty after pushes")
	}
	value, err := s.Pop()
	if err != nil {
		t.Error("unexpected error on Pop:", err)
	}
	if value != 30 {
		t.Errorf("expected popped value 30, got %d", value)
	}
	value, err = s.Pop()
	if err != nil {
		t.Error("unexpected error on Pop:", err)
	}
	if value != 20 {
		t.Errorf("expected popped value 20, got %d", value)
	}
	value, err = s.Pop()
	if err != nil {
		t.Error("unexpected error on Pop:", err)
	}
	if value != 10 {
		t.Errorf("expected popped value 10, got %d", value)
	}
	_, err = s.Pop()
	if err == nil {
		t.Error("expected error when popping from empty stack")
	}
}

func TestPeek(t *testing.T) {
	s := NewStack[string]()
	_, err := s.Peek()
	if err == nil {
		t.Error("expected error on Peek from empty stack")
	}
	s.Push("foo")
	top, err := s.Peek()
	if err != nil {
		t.Error("unexpected error on Peek:", err)
	}
	if top != "foo" {
		t.Errorf("expected Peek value 'foo', got '%s'", top)
	}
	s.Push("bar")
	top, err = s.Peek()
	if err != nil {
		t.Error("unexpected error on Peek:", err)
	}
	if top != "bar" {
		t.Errorf("expected Peek value 'bar', got '%s'", top)
	}
}

func TestClear(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Clear()
	if !s.IsEmpty() {
		t.Error("expected stack to be empty after Clear")
	}
	if s.Size() != 0 {
		t.Errorf("expected size 0 after Clear, got %d", s.Size())
	}
	_, err := s.Pop()
	if err == nil {
		t.Error("expected error when popping from cleared stack")
	}
}

func TestClone(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	clone := s.Clone()
	if clone.Size() != s.Size() {
		t.Errorf("expected clone size %d, got %d", s.Size(), clone.Size())
	}
	clone.Push(3)
	if s.Size() == clone.Size() {
		t.Error("expected original and clone to diverge after modifying clone")
	}
}

func TestReverse(t *testing.T) {
	s := NewStack[int]()
	for i := 1; i <= 3; i++ {
		s.Push(i)
	}
	s.Reverse()
	expected := []int{3, 2, 1}
	actual := s.ToSlice()
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("expected %v at position %d, got %v", v, i, actual[i])
		}
	}
}

func TestToSlice(t *testing.T) {
	s := NewStack[int]()
	s.Push(42)
	s.Push(7)
	slice := s.ToSlice()
	if len(slice) != 2 || slice[0] != 42 || slice[1] != 7 {
		t.Errorf("unexpected slice content: %v", slice)
	}
}

func TestCapacity(t *testing.T) {
	s := NewStack[int]()
	if s.Capacity() != 0 {
		t.Errorf("expected initial capacity 0, got %d", s.Capacity())
	}
	for i := range 100 {
		s.Push(i)
	}
	if s.Capacity() < s.Size() {
		t.Errorf("expected capacity >= size, got capacity %d, size %d", s.Capacity(), s.Size())
	}
}

func TestContains(t *testing.T) {
	s := NewStack[string]()
	s.Push("apple")
	s.Push("banana")
	if !s.Contains("apple") {
		t.Error("expected stack to contain 'apple'")
	}
	if s.Contains("orange") {
		t.Error("did not expect stack to contain 'orange'")
	}
}

func TestString(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	got := s.String()
	want := "Stack: [1 2]"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
