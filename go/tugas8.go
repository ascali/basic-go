package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	rand.Seed(time.Now().Unix())
	var message = make(chan string)
	go sendData(message)
	getData(message)
}

func sendData(ch chan<- string) {
	for i := 0; true; i++ {
		ch <- "Apakah Teman Teman"
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func getData(ch <-chan string) {
loop:
	select {
	case data := <-ch:
		fmt.Println("Receive data :", data)
	case <-time.After(time.Second * 5):
		fmt.Print("timeout for 5 sec...")
		break loop
	}
}
