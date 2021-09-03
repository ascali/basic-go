package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text string
	fmt.Print("Input string : ")
	fmt.Scanln(&text)

	var regex, errRegex = regexp.Compile(`[a-z]+|[A-Z]+|\d`)
	if errRegex != nil {
		fmt.Println(errRegex.Error())
	}
	var resRegex = regex.FindAllString(text, -1)
	fmt.Println(resRegex)
	var isMatchRegex = regex.MatchString(text)
	fmt.Println(isMatchRegex)
	var isIndexRegex = regex.FindStringIndex(text)
	fmt.Println(isIndexRegex)
	var isNewStrRegex = regex.ReplaceAllString(text, "Apple")
	fmt.Println(isNewStrRegex)
}
