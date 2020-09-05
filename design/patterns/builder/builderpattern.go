package main

import (
	"strings"
	"fmt"
)

type shield struct{
	front bool
	back bool
	right bool
	left bool
}

type shBuilder struct{
	code string
}

func NewShieldBuilder() *shBuilder  {
	return new(shBuilder)
}


func (s *shBuilder) RaiseFront() *shBuilder{
	s.code += "F"
	return s
}


func (s *shBuilder) RaiseBack() *shBuilder{
	s.code += "B"
	return s
}


func (s *shBuilder) RaiseRight() *shBuilder{
	s.code += "R"
	return s
}

func (s *shBuilder) RaiseLeft() *shBuilder{
	s.code += "L"
	return s
}

func (s *shBuilder) Build() *shield {
	code := s.code

	return &shield{
		front : strings.Contains(code,"F"),
		back : strings.Contains(code,"B"),
		right : strings.Contains(code,"R"),
		left : strings.Contains(code,"L"),
	}
}

func main(){
	build := NewShieldBuilder()
	shield := build.RaiseBack().RaiseRight().Build()
	fmt.Println(shield)

}