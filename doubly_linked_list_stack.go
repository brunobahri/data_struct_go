package main

import (
	"fmt"
)

// Node represents a node in the doubly linked list
type Node struct {
	value int
	next  *Node
	prev  *Node
}

// Stack represents a stack using a doubly linked list
type Stack struct {
	top  *Node
	size int
}

// Push adds a new element to the top of the stack
func (s *Stack) Push(value int) {
	newNode := &Node{value: value}
	if s.top == nil {
		s.top = newNode
	} else {
		newNode.prev = s.top
		s.top.next = newNode
		s.top = newNode
	}
	s.size++
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	value := s.top.value
	s.top = s.top.prev
	if s.top != nil {
		s.top.next = nil
	}
	s.size--
	return value, true
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() (int, bool) {
	if s.top == nil {
		return 0, false
	}
	return s.top.value, true
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return s.size
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Printf("Stack size: %d\n", stack.Size())
	
	if value, ok := stack.Peek(); ok {
		fmt.Printf("Top element: %d\n", value)
	} else {
		fmt.Println("Stack is empty")
	}

	for stack.Size() > 0 {
		if value, ok := stack.Pop(); ok {
			fmt.Printf("Popped: %d\n", value)
		} else {
			fmt.Println("Stack is empty")
		}
	}

	fmt.Printf("Stack size after popping all elements: %d\n", stack.Size())
}

