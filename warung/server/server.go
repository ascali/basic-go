package server

import (
	"database/sql"
	"log"
	_ "mysql-master"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/warung_lengko")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
