package main

import (
	"fmt"
	"math/rand"
	"time"
)

var mappings = map[string]int{
	"Facebook": 5,
	"Twitter":  10,
	"Google":   15,
}

func findServer(name, server string, chnl chan int) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	chnl <- mappings[name]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	chnl1 := make(chan int)
	chnl2 := make(chan int)

	name := "Google"

	go findServer(name, "Server 1", chnl1)
	go findServer(name, "Server 2", chnl2)

	select {
	case svr := <-chnl1:
		fmt.Println("server-1", svr)
	case svr := <-chnl2:
		fmt.Println("server-2", svr)
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout!")
	}
}
