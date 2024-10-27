package main

import "fmt"

type stack []string

func (s *stack) push(value string) {
	*s = append(*s, value)
}

func (s *stack) isFull(capacity int) bool {
	return len(*s) >= capacity
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) peek() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	return (*s)[len(*s)-1], true
}

func (s *stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top, true
}

func main() {
	var s stack
	capacity := 3

	// Test if the stack is empty
	fmt.Println("Is the stack empty?", s.isEmpty()) // Should print true

	// Push elements onto the stack
	s.push("hello")
	s.push("world")
	s.push("hi!!!!!")
	fmt.Println("Stack after pushes:", s)

	// Check if the stack is full
	fmt.Println("Is the stack full?", s.isFull(capacity)) // Should print true

	// Peek at the top element
	if top, ok := s.peek(); ok {
		fmt.Println("Top element:", top) // "
	} else {
		fmt.Println("Stack is empty")
	}

	// Pop the top element and print it
	if val, ok := s.pop(); ok {
		fmt.Println("Popped value:", val) // Should print "mambo"
	} else {
		fmt.Println("Stack is empty")
	}

	// Print remaining elements in the stack
	fmt.Println("Stack after pop:", s) // Should print [hello world]

	// Pop again and print the remaining stack
	if val, ok := s.pop(); ok {
		fmt.Println("Popped value:", val) // Should print "world"
	} else {
		fmt.Println("Stack is empty")
	}
	fmt.Println("Stack after another pop:", s) // Should print [hello]

	// Pop the last element
	if val, ok := s.pop(); ok {
		fmt.Println("Popped value:", val) // Should print "hello"
	} else {
		fmt.Println("Stack is empty")
	}

	// Check if the stack is empty after popping all elements
	fmt.Println("Is the stack empty after popping all?", s.isEmpty()) // Should print true
}
