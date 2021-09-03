package main

import "fmt"

func main() {
	buah := []string{"apel", "jeruk", "anggur", "melon"}
	buah = append(buah, "pepaya")
	fmt.Println("Jumlah Element", len(buah))
	fmt.Println("Isi Element :", buah)
	i := 0
	for {
		fmt.Println("Element ke - :", i, " adalah ", buah[i])
		i++
		if i == len(buah) {
			break
		}
	}
}
