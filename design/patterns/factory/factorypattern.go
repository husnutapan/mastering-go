package main

import (
	"errors"
	"fmt"
)

type CommandProcess interface {
	Process(val1, val2 int) int
	WriteConsole(result int)
}

const (
	SUM = iota
	SUBSTACT
	MULTIPY
	DIVIDE
)

func CreateProcess(selectionType int) (CommandProcess, error) {
	switch selectionType {
	case SUM:
		return new(Sum), nil
	case SUBSTACT:
		return new(Substract), nil
	default:
		return nil, errors.New("Invalid process type...")
	}
}

type Sum struct {
}

func (s *Sum) Process(val1, val2 int) int {
	return val1 + val2
}

func (s *Sum) WriteConsole(result int) {
	fmt.Println(result)
}

type Substract struct {
}

func (s *Substract) Process(val1, val2 int) int {
	return val1 - val2
}

func (s *Substract) WriteConsole(result int) {
	fmt.Println(result)
}

func main() {
	cp, err := CreateProcess(1)
	if err == nil {
		cp.WriteConsole(cp.Process(10, 20))
	}
}
