package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	defer recoverMe()
	var input string
	fmt.Print("Input number :")
	fmt.Scanln(&input)

	var isRandNumber int = rand.Intn(5-0) + 0
	var number int
	var err error
	number, err = strconv.Atoi(input)

	if number == isRandNumber && err == nil {
		fmt.Println(number, "is a number we want")
	} else {
		fmt.Println(input, "is not number we want!")
		panic(err.Error())
		fmt.Print("show me")
	}
}

func recoverMe() {
	if r := recover(); r != nil {
		fmt.Println("it's not value we want!")
	}
}

// to decide a number from random number given between 1 to 5 by program
