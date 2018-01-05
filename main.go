package main

import (
	"fmt"
	"net/http"
)

func service(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to your Go application running in Docker!")
}

func main() {
	http.HandleFunc("/", service)
	http.ListenAndServe(":8285", nil)
}