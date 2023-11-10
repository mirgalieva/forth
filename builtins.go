package main

import (
	"fmt"
)

func (e *Evaluator) add() error {
	first, second, err := e.stack.Pop2TopValues()
	if err != nil {
		return err
	}
	e.stack.Push(first + second)
	return nil
}

func (e *Evaluator) sub() error {
	first, second, err := e.stack.Pop2TopValues()
	if err != nil {
		return err
	}
	e.stack.Push(second - first)
	return nil
}

func (e *Evaluator) mul() error {
	first, second, err := e.stack.Pop2TopValues()
	if err != nil {
		return err
	}
	e.stack.Push(first * second)
	return nil
}

func (e *Evaluator) div() error {
	first, second, err := e.stack.Pop2TopValues()
	if err != nil {
		return err
	}
	if first == 0 {
		return fmt.Errorf("division by zero")
	}
	e.stack.Push(second / first)
	return nil
}
func (e *Evaluator) dup() error {
	val, err := e.stack.Top()
	if err != nil {
		return err
	}
	e.stack.Push(val)
	return nil
}
func (e *Evaluator) over() error {
	val, err := e.stack.SecondTop()
	if err != nil {
		return err
	}
	e.stack.Push(val)
	return nil
}
func (e *Evaluator) swap() error {
	first, second, err := e.stack.Pop2TopValues()
	if err != nil {
		return err
	}
	e.stack.Push(first)
	e.stack.Push(second)
	return nil
}

func (e *Evaluator) drop() error {
	_, err := e.stack.Pop()
	if err != nil {
		return err
	}
	return nil
}
