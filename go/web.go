package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Ascaliko")
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Hello World")
	})
	http.HandleFunc("/index", index)
	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
