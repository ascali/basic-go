package library

import "fmt"

func thisPrivate() {
	fmt.Println("this is private")
}

func ThisPublic() {
	fmt.Println("this is public")
	thisPrivate()
}

func TampilMahasiswa(nama string, umur int) {
	var mhs = mahasiswa{nama, umur}
	mhs.methodTampilMahasiswa()
}

type mahasiswa struct {
	nama string
	umur int
}

//method
func (m mahasiswa) methodTampilMahasiswa() {
	fmt.Println("Nama :", m.nama)
	fmt.Println("Umur :", m.umur)
}
