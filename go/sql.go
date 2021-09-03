package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

type Pelajar struct {
	id    int
	name  string
	age   int
	class string
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

	rows, err := db.Query("SELECT * FROM pelajar ORDER BY name ASC")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Pelajar
	for rows.Next() {
		var each = Pelajar{}
		var err = rows.Scan(&each.id, &each.name, &each.age, &each.class)
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
		fmt.Println(each.name, each.age, each.class)
	}

	// fmt.Println(result)
}

func create() {
	db, err := connection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO pelajar VALUES (?, ?, ?, ?)", nil, "DEDI", 15, "VII-B")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully created data")

}

func edit() {
	db, err := connection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE pelajar SET age = ? WHERE name = ?", 20, "DEDI")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully update data")
}
func delete() {
	db, err := connection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM pelajar WHERE name = ?", "Godi")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully update data")
}

func main() {
	// create()
	// edit()
	// delete()
	read()
}
