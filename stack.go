package main

import "fmt"

type Stack struct {
	stack []int
}

func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}
func (s *Stack) Push(value int) {
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("empty stack")
	}
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return val, nil
}

func (s *Stack) Pop2TopValues() (int, int, error) {
	first, err := s.Pop()
	if err != nil {
		return 0, 0, err
	}
	second, err := s.Pop()
	if err != nil {
		return 0, 0, err
	}
	return first, second, nil
}

func (s *Stack) GetStack() []int {
	return s.stack
}

func (s *Stack) Top() (int, error) {
	if s.Empty() {
		return 0, fmt.Errorf("empty stack")
	}
	val := s.stack[len(s.stack)-1]
	return val, nil
}

func (s *Stack) SecondTop() (int, error) {
	if len(s.stack) < 2 {
		return 0, fmt.Errorf("stack has less than two arguments")
	}
	return s.stack[len(s.stack)-2], nil
}
