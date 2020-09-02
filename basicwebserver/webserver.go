package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the root endpoint")
}
