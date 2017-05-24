package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, golang!")
}
