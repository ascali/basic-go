package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

type Pegawai struct {
	nik           string
	nama_pegawai  string
	tempat_lahir  string
	tanggal_lahir string
	jenis_kelamin string
	alamat        string
	pendidikan    string
}

func connection() (*sql.DB, error) {
	var db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func read() {
	db, err := connection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM pegawai ORDER BY nama_pegawai ASC")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Pegawai
	for rows.Next() {
		var each = Pegawai{}
		var err = rows.Scan(&each.nik, &each.nama_pegawai, &each.tempat_lahir, &each.tanggal_lahir, &each.jenis_kelamin, &each.alamat, &each.pendidikan)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println("-------------------------")
		fmt.Println("nik", each.nik, "\nnama_pegawai", each.nama_pegawai, "\ntempat_lahir", each.tempat_lahir, "\ntanggal_lahir", each.tanggal_lahir, "\njenis_kelamin", each.jenis_kelamin, "\nalamat", each.alamat, "\npendidikan", each.pendidikan)
		fmt.Println("-------------------------")
	}

	// fmt.Println(result)
}

func main() {
	read()
}
