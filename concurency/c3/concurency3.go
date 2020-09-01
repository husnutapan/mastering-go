package main

import (
	"fmt"
	"sync"
	"time"
)

type ThreadSafe struct {
	value map[string]int
	mux   sync.Mutex
}

func (t *ThreadSafe) Increment(key string) {
	t.mux.Lock()
	t.value[key]++
	t.mux.Unlock()
}

func (t *ThreadSafe) Get(key string) int {
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.value[key]
}

func main() {

	c := ThreadSafe{value: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Increment("myKey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Get("myKey"))

}
