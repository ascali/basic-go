package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "mysql-master"
	"net/http"
)

type FoodMenu struct {
	ID       string
	MenuName string
	Price    int
}

var data = []FoodMenu{
	{"M01", "Rumbah Kucur", 5000},
	{"M02", "Rumbah Uleg", 6000},
	{"M02", "Rumbah Asem", 7000},
}

// type pelajar_struct struct {
// 	id    int
// 	name  string
// 	age   int
// 	class string
// }

type pelajar_struct struct {
	id    int    `json:"id"`
	name  string `json:"name"`
	age   int    `json:"age"`
	class string `json:"class"`
}

var pelajarData []pelajar_struct

func main() {
	read_pelajar()

	http.HandleFunc("/pelajar", GetPelajar)
	http.HandleFunc("/menu", GetMenu)
	http.HandleFunc("/search_menu", SearchMenu)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{"status": "true", "message": "Welcome to web api"}
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			var result, err = json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(result)
			return
		}
		http.Error(w, "", http.StatusBadRequest)
	})

	fmt.Println("Running on 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func SearchMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var IsMenuName = r.FormValue("MenuName")
		var result []byte
		var err error

		for _, each := range data {
			if each.MenuName == IsMenuName {
				result, err = json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(result)
				return
			}
		}
		http.Error(w, "Menu not available", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func GetPelajar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var res, err = json.Marshal(pelajarData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func connection() (*sql.DB, error) {
	var db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func read_pelajar() {
	db, err := connection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	pelajar, err := db.Query("SELECT * FROM pelajar")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer pelajar.Close()

	for pelajar.Next() {
		var each_pelajar pelajar_struct
		var err = pelajar.Scan(&each_pelajar.id, &each_pelajar.name, &each_pelajar.age, &each_pelajar.class)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		pelajarData = append(pelajarData, each_pelajar)
	}

	if err = pelajar.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
