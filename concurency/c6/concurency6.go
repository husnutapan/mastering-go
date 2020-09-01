package main

import (
	"fmt"
	"time"
)

func main() {
	chn := make(chan string, 2)

	go sayMultipleTime(chn, 5)

	for s := range chn {
		fmt.Println(s)
	}

	v, ok := <-chn
	fmt.Println(!ok, v)
}

func sayMultipleTime(chn chan string, k int) {
	for i := 0; i <= k; i++ {
		chn <- "Hello"
		time.Sleep(time.Second)
		fmt.Println("Break")
	}

	close(chn)
}
