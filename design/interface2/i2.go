package main

import "fmt"

func main() {
	printType("text")
	printType(1)
	printType(1.0)
}

func printType(i interface{}) {
	switch i := i.(type) {
	case string:
		fmt.Println("String type ", i)
	case int:
		fmt.Println("Integer type ", i)
	case float64:
		fmt.Println("Float type ", i)
	}
}
