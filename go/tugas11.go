package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 string
	var num2 string
	fmt.Print("Input Number1 :")
	fmt.Scanln(&num1)
	fmt.Print("Input Number2 :")
	fmt.Scanln(&num2)
	var isNum1, errNum1 = strconv.Atoi(num1)
	var isNum2, errNum2 = strconv.Atoi(num2)
	if errNum1 == nil && errNum2 == nil {
		fmt.Println(isNum1, "+", isNum2, "=", (isNum1 + isNum2))
		fmt.Println(isNum1, "-", isNum2, "=", (isNum1 - isNum2))
		fmt.Println(isNum1, "*", isNum2, "=", (isNum1 * isNum2))
		fmt.Println(isNum1, "/", isNum2, "=", (float64(isNum1) / float64(isNum2)))
	}
}
