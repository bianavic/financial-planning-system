package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, Stranger! Welcome to my Home Page")
	})

	http.ListenAndServe(":8888", nil)
}
