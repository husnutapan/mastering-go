package main

import (
	"sync"
	"fmt"
)

func onceBody(){
	fmt.Println("Oncebody function called from out")
}



func main(){
	 var once sync.Once 
	 done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func ()  {
		once.Do(onceBody)
		done <- true
	}()
	}

	for i := 0; i < 10; i++ {
		<- done
	}
}