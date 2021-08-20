package main

import (
	"fmt"
)

//NodeStack represents a stack node
type NodeStack struct {
	Value int
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*NodeStack
	len   int
}

// Push adds a node to the stack.
func (s *Stack) Push(value int) {
	n := &NodeStack{value}
	s.nodes = append(s.nodes, n)
	s.len++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *NodeStack {
	if s.len == 0 {
		return nil
	}
	s.len--
	popped := s.nodes[s.len]
	s.nodes = s.nodes[:s.len]
	return popped
}

// Peek returns the last element in the stack
func (s *Stack) Peek() *NodeStack {
	return s.nodes[s.len-1]
}

// IsEmpty returns if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func main() {
	mystack := NewStack()
	fmt.Println("IsEmpty : ", mystack.IsEmpty())

	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)

	fmt.Println("len : ", mystack.len)
	fmt.Println("popped : ", mystack.Pop())
	fmt.Println("len : ", mystack.len)

	mystack.Push(4)
	mystack.Push(5)
	fmt.Println("len : ", mystack.len)

	fmt.Println("peek : ", mystack.Peek())

	fmt.Println("IsEmpty : ", mystack.IsEmpty())
}
