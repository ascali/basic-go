package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("My Name is Ascaliko")
	TimeSleep()
	fmt.Println("My Hobby is coding")
	TimeSleep()
	fmt.Println("My D.o.B is 15 July 1993")
}

func TimeSleep() {
	time.Sleep(time.Second * 5)
}
