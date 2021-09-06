package stack

import (
	// "errors"
	"fmt"
)

type Stack struct {
	data []int
	size int
}

func NewStack(cap int) *Stack {
	return &Stack{ data: make([]int, 0, cap), size: 0 }
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Pop() interface{}{
	if s.size <= 0 {
		return nil
	}
	n := s.data[s.size-1]
	s.data = s.data[:s.size-1]
	s.size--
	return n
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
	s.size++
}

func (s *Stack) GetTop() interface{}{
	if s.size <= 0 {
		return nil
	}
	return s.data[s.size-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
    return s.size == 0
}

// String implements Stringer interface
func (s *Stack) String() string {
    return fmt.Sprint(s.data)
}