package main

import (
	"errors"
	"fmt"
)

var InvalidNode = errors.New("Node is not valid")

type Node interface {
	SetValue(v int) error
	GetValue() int
}
type LinkedList struct {
	next    *LinkedList
	value   int
	message string
}

func (node *LinkedList) SetValue(value int) error {
	if node == nil {
		return InvalidNode
	}
	node.value = value
	return nil
}
func NewLinkedList() *LinkedList {
	return &LinkedList{message: "normal node message test"}
}

func (node *LinkedList) GetValue() int {
	return node.value
}

type PowerNode struct {
	next    *PowerNode
	value   int
	message string
}

func (node *PowerNode) SetValue(value int) error {
	if node == nil {
		return InvalidNode
	}
	node.value = value * 10
	return nil
}

func (node *PowerNode) GetValue() int {
	return node.value
}

func NewPowerNode() *PowerNode {
	return &PowerNode{message: "power node message test"}
}

func createNode(value int) Node {
	normalNode := NewLinkedList()
	normalNode.SetValue(value)
	return normalNode
}

func main() {
	var n Node
	fmt.Println(n.SetValue(3))
	/*node := createNode(5)

	switch n := node.(type) {
	case *LinkedList:
		fmt.Println("Normal Node detected ", n.message)
	case *PowerNode:
		fmt.Println("Power Node detected ", n.message)
	}*/
}

