package main

import "fmt"

func main() {
	bufferedChan := make(chan string, 2)

	bufferedChan <- "Text1"
	bufferedChan <- "Text2"

	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)
}
