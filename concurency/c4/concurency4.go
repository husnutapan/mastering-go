package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)

	go waitAndSay(channel, "Var1")
	time.Sleep(time.Second)
	channel <- true

	<-channel
}

func waitAndSay(channel chan bool, text string) {
	if a := <-channel; a {
		fmt.Println(text)
	}
	time.Sleep(2 * time.Second)

	channel <- true
}
