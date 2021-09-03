package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	var input string
	fmt.Print("Input text to encode to base64 & SHA1 :")
	fmt.Scanln(&input)

	fmt.Println("-------------------------------")

	var encodeStr = base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println("Encode :", encodeStr)
	var decodeStr, _ = base64.StdEncoding.DecodeString(encodeStr)
	fmt.Println("Decode :", string(decodeStr))

	fmt.Println("-------------------------------")

	var sha = sha1.New()
	sha.Write([]byte(input))
	var encryptionSha = sha.Sum(nil)
	var encryptionShaStr = fmt.Sprintf("%x", encryptionSha)
	fmt.Println("SHA1 :", encryptionShaStr)

	fmt.Println("-------------------------------")
}
