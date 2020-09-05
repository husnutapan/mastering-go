package main

import (
	"net/http"
	"log"
	"os"
	"sync"
	"fmt"
)

var mylogger *myCustomLogger
var once sync.Once

type myCustomLogger struct {
	*log.Logger
	filename string
}

func Instance() *myCustomLogger {
	once.Do(func() {
		fmt.Println("Instance adding to context")
		mylogger = createLogger("mycustomlogger.log")
	})
	return mylogger
}

func createLogger(filename string) *myCustomLogger {
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	return &myCustomLogger{
		Logger:   log.New(file, " custom ", log.Lshortfile),
		filename: filename,
	}
}
func main() {
	logger := Instance()
	logger.Println("Starting web service")


	http.HandleFunc("/",rootPath)
	http.ListenAndServe(":8080",nil)
}

func rootPath(w http.ResponseWriter,r *http.Request)  {
	logger:= Instance()
	fmt.Fprint(w,"Welcome base root path")
	logger.Println("Recieved http get request")
}
