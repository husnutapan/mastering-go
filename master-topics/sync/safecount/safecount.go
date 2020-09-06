package main

import (
	"sync" 
	"fmt"
)

type safecounter struct{
	count int
	sync.Mutex
}


func (sc *safecounter) Increment(){
	sc.Lock()
	sc.count++
	sc.Unlock()
}

func (sc *safecounter) Decrement(){
	sc.Lock()
	sc.count--
	sc.Unlock()
}

func (sc *safecounter) GetCounter() int{
	//prevent race condition...
	sc.Lock()
	result:=sc.count
	sc.Unlock()
	return result
}


func main(){
	sc := new(safecounter)
	for i := 0; i < 1000; i++ {
		go sc.Increment()
		go sc.Decrement()
	}
	fmt.Println(sc.GetCounter())
}