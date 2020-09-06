package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

type mapStruct struct{
	myMap map[int]int
	sync.RWMutex
}

func runWriters(ms *mapStruct,loopSize int){
	for i := 0; i < loopSize; i++ {
		ms.Lock()
		ms.myMap[i]=1*10
		ms.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func runReaders(ms *mapStruct,limit int, threadname string){
	for {
		ms.RLock()
		r := ms.myMap[rand.Intn(limit)]
		ms.RUnlock()
		fmt.Println(r)
		time.Sleep(1 * time.Second)
	}
}
func main(){
	ms := &mapStruct{myMap:make(map[int]int)}
	go runWriters(ms,10)
	go runReaders(ms,10,"thread-1")
	go runReaders(ms,10,"thread-2")

	time.Sleep(15 * time.Second)
}