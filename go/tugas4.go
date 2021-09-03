package main

import "fmt"

func main() {
	mahasiswa := map[string]int{"Aldo": 182, "Yosep": 178}
	tampil_mahasiswa(mahasiswa)
}

func tampil_mahasiswa(mahasiswa map[string]int) {
	fmt.Println("Aldo :", mahasiswa["Aldo"], "cm")
	fmt.Println("Yosep :", mahasiswa["Yosep"], "cm")
}
