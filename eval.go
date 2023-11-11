//go:build !solution
// +build !solution

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const endOfCommand = ";"

type Command struct {
	ID              int
	Function        func() error
	Definition      []int
	NumberToPush    int
	hasNumberToPush bool
}

type Evaluator struct {
	stack          Stack
	commandsData   map[string]Command
	commandsList   map[int]Command
	commandsNumber int
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	e := Evaluator{}
	e.commandsNumber = 8
	add := Command{Function: e.add, ID: 0}
	sub := Command{Function: e.sub, ID: 1}
	mul := Command{Function: e.mul, ID: 2}
	div := Command{Function: e.div, ID: 3}
	over := Command{Function: e.over, ID: 4}
	dup := Command{Function: e.dup, ID: 5}
	drop := Command{Function: e.drop, ID: 6}
	swap := Command{Function: e.swap, ID: 7}
	e.commandsData = map[string]Command{
		"+":    add,
		"-":    sub,
		"*":    mul,
		"/":    div,
		"over": over,
		"dup":  dup,
		"drop": drop,
		"swap": swap,
	}
	e.commandsList = map[int]Command{0: add, 1: sub, 2: mul, 3: div, 4: over, 5: dup, 6: drop, 7: swap}
	return &e
}

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.
func (e *Evaluator) Process(row string) ([]int, error) {
	if len(row) == 0 {
		return make([]int, 0), fmt.Errorf("empty command")
	}
	row = strings.ToLower(row)
	if row[0] == ':' {
		return e.reserveNewCommand(row)
	}
	data := strings.Split(row, " ")
	err := e.parseSentence(data)
	if err != nil {
		return make([]int, 0), err
	}
	return e.stack.GetStack(), nil
}

func (e *Evaluator) parseSentence(data []string) error {
	for _, word := range data {
		// if number, push to the stack
		if number, err := strconv.Atoi(word); err == nil {
			e.stack.Push(number)
			continue
		}
		// check if word is reserved command
		c, ok := e.commandsData[word]
		if !ok {
			return fmt.Errorf("undefined command")
		}
		ok, err := e.evaluate(c)
		if err != nil {
			return err
		}
		if ok {
			continue
		}
		// if command is not builtin, we sequentially evaluate commands from its definition
		for _, command := range c.Definition {
			commandData, ok := e.commandsList[command]
			if !ok {
				return fmt.Errorf("undefined command")
			}
			_, err := e.evaluate(commandData)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *Evaluator) reserveNewCommand(row string) ([]int, error) {
	data := strings.SplitN(row[1:], " ", 3)
	if len(data) < 3 {
		return make([]int, 0), fmt.Errorf("incorrect command definition")
	}
	Name, commandStr := data[1], data[2]
	if _, err := strconv.Atoi(Name); err == nil {
		return make([]int, 0), fmt.Errorf("can not redefine number")
	}
	definition, err := e.makeDefinition(strings.Split(commandStr, " "))
	if err != nil {
		return make([]int, 0), err
	}
	// if it is a redefinition, we should delete previous definition
	if _, ok := e.commandsData[Name]; ok {
		delete(e.commandsData, Name)
	}
	e.addNewCommand(Name, Command{Definition: definition, ID: e.commandsNumber})
	return e.stack.GetStack(), nil
}

func (e *Evaluator) addNewCommand(commandName string, command Command) {
	e.commandsList[e.commandsNumber] = command
	e.commandsData[commandName] = command
	e.commandsNumber++
}

func (e *Evaluator) evaluate(command Command) (bool, error) {
	if command.Function != nil {
		err := command.Function
		if err() != nil {
			return false, err()
		}
		return true, nil
	} else if command.hasNumberToPush {
		e.stack.Push(command.NumberToPush)
		return true, nil
	}
	return false, nil
}

func (e *Evaluator) makeDefinition(commands []string) ([]int, error) {
	definition := make([]int, 0)
	for _, command := range commands {
		if command == endOfCommand {
			break
		}
		if c, ok := e.commandsData[command]; ok {
			if c.Function != nil {
				definition = append(definition, c.ID) // if it's builtin function
			} else {
				definition = append(definition, c.Definition...) // if it was defined earlier
			}
		} else if number, err := strconv.Atoi(command); err == nil { //if it's pushing a number
			definition = append(definition, e.commandsNumber)
			e.addNewCommand(command, Command{NumberToPush: number, hasNumberToPush: true, ID: e.commandsNumber})
		} else {
			return make([]int, 0), fmt.Errorf("undefined command") // else we don't kniw this command
		}
	}
	return definition, nil
}
