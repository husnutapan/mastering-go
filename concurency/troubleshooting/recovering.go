package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findSrv1(name, server string) (int, error) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	if i, ok := mappings[name]; !ok {
		panic("Member not found...")
	} else {
		return i, nil
	}
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovering panic...", err)
		}
	}()

	if _, err := findSrv1("Amazon", "Server 1"); err != nil {
		fmt.Println("Error occured while searching data", err)
	}
}
