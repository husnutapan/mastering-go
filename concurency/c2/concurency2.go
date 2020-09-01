package main

import (
	"fmt"
	"time"
)

func main() {
	word := "var1"

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(word)
	}()
	fmt.Println("Var2")
	word = "var1 updated"
	time.Sleep(3 * time.Second)
}
