package stack

import (
	// "errors"
	// "fmt"
)

type Stack struct {
	top *Node
	length int
}

type Node struct {
	val interface{}
	prev *Node
}

func NewStack() *Stack {
	return &Stack{ nil, 0 }
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Pop() interface{}{
	if s.length <= 0 {
		return nil
	}
	elem := s.top
	s.top = elem.prev
	s.length--
	return elem.val
}

func (s *Stack) Push(val interface{}) {
	elem := &Node{ val, s.top }
	s.top = elem
	s.length++
}

func (s *Stack) GetTop() interface{}{
	if s.length <= 0 {
		return nil
	}
	return s.top.val
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
    return s.length == 0
}