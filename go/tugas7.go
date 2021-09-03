package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(16)
	go bacatipeInt()
	bacatipeString()
	var input string
	fmt.Scanln(&input)
}

func bacatipeInt() {
	var bacatipe1 int = 1
	var reflectNumber = reflect.ValueOf(bacatipe1)
	fmt.Println("Type", reflectNumber.Type())
	if reflectNumber.Kind() == reflect.Int {
		fmt.Println("Value", reflectNumber.Int())
	}
}

func bacatipeString() {
	var bacatipe2 string = "string"
	var reflectString = reflect.ValueOf(bacatipe2)
	fmt.Println("Type", reflectString.Type())
	if reflectString.Kind() == reflect.String {
		fmt.Println("Value", reflectString.String())
	}
}
