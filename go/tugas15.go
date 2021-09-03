package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var i int = 0
	for {
		fmt.Fprintln(w, i+1)
		i++
		if i == 100 {
			break
		}
	}
}

func utama(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `Hello Ascaliko go to index`)
}

func main() {
	http.HandleFunc("/", utama)
	http.HandleFunc("/index", index)
	fmt.Println("server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
