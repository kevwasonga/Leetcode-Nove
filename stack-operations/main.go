package main

import (
	"fmt"
	"unicode/utf8"
)

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

// compareRunes compares two strings based on the first rune.
func compareRunes(a, b string) bool {
	runeA, _ := utf8.DecodeRuneInString(a)
	runeB, _ := utf8.DecodeRuneInString(b)
	return runeA < runeB
}

func sortStack(stack1, stack2 *stack) {
	for !stack1.isEmpty() {
		value, _ := stack1.pop()
		for !stack2.isEmpty() && !compareRunes(value, (*stack2)[len(*stack2)-1]) {
			temp, _ := stack2.pop()
			stack1.push(temp)
		}
		stack2.push(value)
	}

	for !stack2.isEmpty() {
		value, _ := stack2.pop()
		stack1.push(value)
	}
}

func main() {
	var s stack
	capacity := 5

	fmt.Println("Is the stack empty?", s.isEmpty())

	s.push("banana")
	s.push("apple")
	s.push("cherry")
	s.push("date")
	s.push("elderberry")

	fmt.Println("Stack after pushes:", s)

	fmt.Println("Is the stack full?", s.isFull(capacity))

	if top, ok := s.peek(); ok {
		fmt.Println("Top element:", top)
	} else {
		fmt.Println("Stack is empty")
	}

	sortStack(&s, &stack{})

	fmt.Println("Sorted stack:", s)

	if val, ok := s.pop(); ok {
		fmt.Println("Popped value:", val)
	} else {
		fmt.Println("Stack is empty")
	}

	fmt.Println("Stack after pop:", s)
}
