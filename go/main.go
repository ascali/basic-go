package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	library "lib"
	"math/rand"
	"net/url"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var msg = make(chan string)

func main() {
	isLine := "\n --------------------------- \n"
	fmt.Println("Hello it's me")
	// comment in line
	/* comment
	multi line */
	var fName string = "Asca"
	var lName = "Liko"
	mName := "bin"
	fmt.Println("My name is", fName, mName, lName)
	fmt.Println(isLine)

	// data type
	var isUint uint8 = 255 //0-255
	var isInt int = -128   //-128-127
	var decimal float32 = 3.14
	var isBoolean = true
	var isString string = "This is string"
	fmt.Println(" uint8 = ", isUint, "\n int = ", isInt, " \n decimal = ", decimal, " \n boolean = ", isBoolean, " \n string = ", isString)
	fmt.Println(isLine)

	// constant
	const pi = 3.14
	fmt.Printf("is Pi %.2f", pi)
	fmt.Println(isLine)

	// operator
	var x = 10
	var y = 5
	fmt.Printf("+ %d \n", x+y)
	fmt.Printf("- %d \n", x-y)
	fmt.Printf("* %d \n", x*y)
	fmt.Printf("/ %d \n", x/y)
	fmt.Printf("modulus %d \n", x%y)
	fmt.Println(isLine)

	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)
	fmt.Println(isLine)

	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)
	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)
	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
	fmt.Println(isLine)

	// condition if else & switch
	if x == 10 {
		fmt.Println("Pass Perfect")
	} else if x >= 5 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	fmt.Println(isLine)

	switch y {
	case 10:
		fmt.Println("Pass Perfect")
	case 5:
		fmt.Println("Pass")
	default:
		fmt.Println("Fail")
	}
	fmt.Println(isLine)

	// looping for
	for i := 0; i <= x; i++ {
		fmt.Printf("looping from %d \n", i)
	}
	// diff way
	i := 0
	for {
		fmt.Printf("looping from %d \n", i)
		i++
		if i == 10 {
			break
		}
	}
	fmt.Println(isLine)

	// array
	fruits := []string{"apple", "grape", "banana"}
	fruits[1] = "apple rose"
	fmt.Println("what is it ?", fruits)
	fmt.Println("total ", len(fruits))
	fmt.Println(isLine)

	// slice
	fruits = append(fruits, "mango")
	fmt.Println("add new", fruits)
	fmt.Println("total ", len(fruits))
	fmt.Println(isLine)

	// map
	food_price := map[string]int{"burger": 2, "pizza": 3}
	fmt.Printf("food price for burger is $%d", food_price["burger"])
	fmt.Println(isLine)

	// call function
	message := "this is message from main function"
	fmt.Println(test(message))
	fmt.Println(isLine)
	// function multiple return
	x0, x1, x2, x3, x4 := isMessage(message, x, y, food_price)
	fmt.Println(x0)
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(isLine)

	// variadic
	avg := isCount(1, 2, 3, 4, 5, 6, 7, 3, 3, 3, 3, 3)
	isAvg := fmt.Sprintf("Average %.2f", avg)
	fmt.Println(isAvg)
	fmt.Println(isLine)

	//pointer
	var no1 int = 11
	var addressOfNo1 *int = &no1
	var str1 string = "is string"
	var addressOfStr1 *string = &str1
	fmt.Println("value of no1 :", no1)
	fmt.Println("address of no1 :", addressOfNo1)
	fmt.Println("value of str1 :", str1)
	fmt.Println("address of addressOfStr1 :", addressOfStr1)
	fmt.Println(isLine)

	// struct
	var isStudent student
	isStudent.name = "asca"
	isStudent.class = 12
	isStudent.age = 18
	fmt.Println("Name :", isStudent.name)
	fmt.Println("Class :", isStudent.class)
	fmt.Println("Age :", isStudent.age)
	fmt.Println(isLine)

	// method
	var s2 = student{"Ascaliko", 11, 17}
	s2.myName()
	fmt.Println(isLine)

	// public n private
	library.ThisPublic()
	fmt.Println(isLine)

	// reflect
	var number = 10
	var reflectNumber = reflect.ValueOf(number)
	fmt.Println("Type", reflectNumber.Type())
	if reflectNumber.Kind() == reflect.Int {
		fmt.Println("Type", reflectNumber.Int())
	}
	fmt.Println(isLine)

	// goroutine async func / does'nt wait for completing
	runtime.GOMAXPROCS(1)
	go showMessage(7, "Send message first")
	showMessage(7, "Send message second")
	var input string
	fmt.Scanln(&input)
	fmt.Println(isLine)

	// channel to connect another goroutine
	runtime.GOMAXPROCS(1)
	go messageChannel("a")
	go messageChannel("b")
	go messageChannel("c")
	var message1 = <-msg
	fmt.Println(message1)
	var message2 = <-msg
	fmt.Println(message2)
	var message3 = <-msg
	fmt.Println(message3)
	fmt.Println(isLine)

	// buffered channel
	runtime.GOMAXPROCS(1)
	var messages = make(chan int, 2)
	go func() {
		for {
			i := <-messages
			fmt.Println("Receive data", i)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("Send data", i)
		messages <- i
	}
	fmt.Println(isLine)

	// channel select
	runtime.GOMAXPROCS(1)
	var numberSlice = []int{0, 3, 7, 5, 9, 1}
	fmt.Println("Number", numberSlice)

	var channel1 = make(chan float64)
	go getAverage(numberSlice, channel1)
	var channel2 = make(chan int)
	go getMaximal(numberSlice, channel2)

	for i := 0; i < 2; i++ {
		select {
		case average := <-channel1:
			fmt.Printf("average \t: %.2f \n", average)
		case maximal := <-channel2:
			fmt.Printf("maximal \t: %d \n", maximal)
		}
	}
	fmt.Println(isLine)

	// channel timeout
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)
	var messageT = make(chan int)
	go sendData(messageT)
	getData(messageT)
	fmt.Println(isLine)

	// defer & exit
	defer fmt.Println("defer 1 Hello")
	fmt.Println("welcome")
	fmt.Println("to my house")
	test("What's up bro!")
	fmt.Println(isLine)
	// os.Exit(1)

	// error, panic & recover
	defer recoverMe()
	var inputU string
	fmt.Print("Input number :")
	fmt.Scanln(&inputU)

	var num1 int
	var err error
	num1, err = strconv.Atoi(inputU)
	if err == nil {
		fmt.Println(num1, "is a number")
	} else {
		fmt.Println(num1, "is not number")
		panic(err.Error())
		fmt.Print("show me")
	}
	fmt.Println(isLine)

	// timer
	fmt.Println("starting...")
	time.Sleep(time.Second * 3)
	fmt.Println("done after 3 sec...")
	fmt.Println(isLine)

	// data type convertion
	var str2 = "156"
	var numStr2, errStr2 = strconv.Atoi(str2)
	var num2 = 172
	var strNum2 = strconv.Itoa(num2)
	var flt2 = "3.14"
	var fltStr2, errFlt2 = strconv.ParseFloat(flt2, 64)
	if errStr2 == nil {
		fmt.Println(numStr2)
		fmt.Println(numStr2 + 11)
	}
	fmt.Println(strNum2 + " its number!")
	if errFlt2 == nil {
		fmt.Println(fltStr2)
	}
	fmt.Println(isLine)

	// string function
	var thereIsStrings = strings.Contains("There is strings here ?", "ring")
	var thereIsStrings2 = strings.HasPrefix("There is strings here ?", "The")
	var thereIsStrings3 = strings.HasSuffix("There is strings here ?", "?")
	var thereIsStringsCount = strings.Count("There is strings here ?", "e")
	var thereIsStringsIndex = strings.Index("There is strings here ?", "h")
	var stringReplace = strings.Replace("Hello world", "o", "i", 2)
	fmt.Println(thereIsStrings)
	fmt.Println(thereIsStrings2)
	fmt.Println(thereIsStrings3)
	fmt.Println("How much alphabet 'e'", thereIsStringsCount)
	fmt.Println("Index of alphabet 'h'", thereIsStringsIndex)
	fmt.Println("Hello world replace o to i", stringReplace)
	fmt.Println(isLine)

	// regex
	var text = "orange1 GRAPE banana"
	var regex, errReg = regexp.Compile(`[a-z]+\d`)
	if errReg != nil {
		fmt.Println(errReg.Error())
	}
	var resRegex = regex.FindAllString(text, -1) // -1 is to get all string
	fmt.Println(resRegex)
	var isMatchRegex = regex.MatchString(text)
	fmt.Println(isMatchRegex)
	var isIndexRegex = regex.FindStringIndex(text)
	fmt.Println(isIndexRegex)
	var isNewStrRegex = regex.ReplaceAllString(text, "Apple")
	fmt.Println(isNewStrRegex)
	fmt.Println(isLine)

	// Encode & Decode Base 64 two encryption
	var myNameIs = "My Name is Ascaliko"
	var encodeStr = base64.StdEncoding.EncodeToString([]byte(myNameIs))
	fmt.Println(encodeStr)
	var decodeStr, _ = base64.StdEncoding.DecodeString(encodeStr)
	fmt.Println(string(decodeStr))
	fmt.Println(isLine)

	// Hash SHA 1 oneway encryption
	var sha = sha1.New()
	sha.Write([]byte(myNameIs))
	var encryptionSha = sha.Sum(nil)
	var encryptionShaStr = fmt.Sprintf("%x", encryptionSha)
	fmt.Println(encryptionShaStr)
	fmt.Println(isLine)

	//exec
	var execTerminal, _ = exec.Command("ls").Output()
	fmt.Println(string(execTerminal))
	fmt.Println(isLine)

	// url parser

	var urlText = "http://google.com:8080/hai?search-book=Lord of The Ring&author=Asca"
	var u, e = url.Parse(urlText)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Println(urlText)
	fmt.Println("host :", u.Host)
	fmt.Println("path :", u.Path)
	var book = u.Query()["search-book"][0]
	fmt.Println("search-book :", book)
	var author = u.Query()["author"][0]
	fmt.Println("author :", author)
	fmt.Println(isLine)

	//json
	var jsonStr = `{"Name": "asca", "Age": 17, "Class": "1A"}`
	var jsonData = []byte(jsonStr)
	var data User
	var errJson = json.Unmarshal(jsonData, &data)
	if errJson != nil {
		fmt.Println(errJson.Error())
		return
	}
	fmt.Println("Name :", data.Name)
	fmt.Println("Age :", data.Age)
	fmt.Println("Class :", data.Class)
	fmt.Println(isLine)

	var obj = []User{{"Asca", 17, "1A"}, {"Liko", 18, "1B"}}
	var jsonDataObj, errObj = json.Marshal(obj)
	if errObj != nil {
		fmt.Println(errObj.Error())
	}
	var jsonString = string(jsonDataObj)
	fmt.Println(jsonString)
	fmt.Println(isLine)
}

type User struct {
	Name  string `json:"Name"`
	Age   int
	Class string `json:"Class"`
}

// function
func test(message string) string {
	return message
}

// function multiple return
func isMessage(message string, x int, y int, isMap map[string]int) (string, int, int, int, map[string]int) {
	var z int = x + y
	var z2 int = x - y
	var z3 int = x * y
	message = message + " " + strconv.Itoa(z)
	return message, z, z2, z3, isMap
}

//variadic
func isCount(n ...int) float64 {
	var t int = 0
	for _, n := range n {
		t += n
		fmt.Println(n)
	}
	r := float64(t) / float64(len(n))
	return r
}

// struct
type student struct {
	name  string
	class int
	age   int
}

//method
func (s student) myName() {
	fmt.Println("Encapsulation Name :", s.name)
	fmt.Println("Encapsulation Class :", s.class)
	fmt.Println("Encapsulation Age :", s.age)
}

// function for goroutine
func showMessage(x int, message string) {
	for i := 0; i < x; i++ {
		fmt.Println((1 + i), message)
	}
}

//func for channel
func messageChannel(value string) {
	var data = fmt.Sprintf("Message %s", value)
	msg <- data
}

// func for select channel
func getAverage(number []int, ch chan float64) {
	var sum = 0
	for _, v := range number {
		sum += v
	}
	ch <- float64(sum) / float64(len(number))
}

func getMaximal(number []int, ch chan int) {
	var max = number[0]
	for _, v := range number {
		if max < v {
			max = v
		}
	}
	ch <- max
}

// func channel timeout
func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func getData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Println("Receive Data :", data)
		case <-time.After(time.Second * 2):
			fmt.Println("Timeout... no activity for 2 second")
			break loop
		}
	}
}

// func error, panic & recover
func recoverMe() {
	if r := recover(); r != nil {
		fmt.Println("Finally it's showing me!")
	}
}
