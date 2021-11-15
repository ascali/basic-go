package server

import (
	"database/sql"
	"log"
	_ "mysql-master"
)

func Koneksi() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_wrg")
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
