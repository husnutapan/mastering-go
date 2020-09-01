package main

import (
	"fmt"
	"time"
)

func waitAndSay(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}

func main() {
	go waitAndSay("Var1")
	fmt.Println("Var2")
	time.Sleep(3 * time.Second)
}
