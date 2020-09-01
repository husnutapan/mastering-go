package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var mappings = map[string]int{
	"Facebook": 5,
	"Twitter":  10,
	"Google":   15,
}

type findError struct {
	Name, Server, Message string
}

var ErrorMemberNotFound = errors.New("Member not found")

func (e findError) Error() string {
	return e.Message
}

func findSrv(name, server string) (int, error) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	if i, ok := mappings[name]; !ok {
		return -1, ErrorMemberNotFound
	} else {
		return i, nil
	}
}
func main_deprecate() {

	_, err := findSrv("Amazon", "Server 1")

	if err == ErrorMemberNotFound {
		fmt.Println("Handled error member not found")
	}

	/*fmt.Println("Result is ", result, " error message ", err)*/

	/*defer func() {
		if err := recover(); err != nil {
			fmt.Println("A panic recovered", err)
		}
	}()*/
}
