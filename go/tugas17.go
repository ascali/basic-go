package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	_ "mysql-master"
	"net/http"
)

type userStruct struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

type Response struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    []userStruct `json:"data"`
}

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var user userStruct
	var arr_user []userStruct
	var response Response
	var isId = r.FormValue("id")

	db := connect()
	defer db.Close()

	if isId != "" {
		row := db.QueryRow("SELECT * FROM users WHERE id = ?", isId).Scan(&user.Id, &user.Name, &user.Age)
		if row != nil {
			log.Fatal(row.Error())
		}
		arr_user = append(arr_user, user)
	} else {
		rows, err := db.Query("SELECT * FROM users ORDER BY id ASC")
		if err != nil {
			log.Print(err)
		}
		for rows.Next() {
			id := &user.Id
			name := &user.Name
			age := &user.Age
			var err = rows.Scan(id, name, age)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				arr_user = append(arr_user, user)
			}
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var arr_user []userStruct
	var response Response

	if r.Method == "POST" {
		var name = r.FormValue("name")
		var age = r.FormValue("age")

		if age != "" || name != "" {
			_, err := db.Exec("INSERT INTO users VALUES (?, ?, ?)", nil, string(name), age)
			if err != nil {
				log.Print(err)
			}

			fmt.Println("Successfully created data")

			response.Status = 1
			response.Message = "Successfully created data"
		} else {
			response.Status = 0
			response.Message = "All field is required"
		}
	} else {
		response.Status = 0
		response.Message = "Unknown method"
	}
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	var arr_user []userStruct
	var response Response

	if r.Method == "PUT" {
		var id = r.FormValue("id")
		var name = r.FormValue("name")
		var age = r.FormValue("age")
		if id != "" || age != "" || name != "" {
			_, err := db.Exec("UPDATE users SET age = ?, name = ? WHERE id = ?", age, string(name), id)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("Successfully update data")

			response.Status = 1
			response.Message = "Successfully update data"
		} else {
			response.Status = 0
			response.Message = "All field is required"
		}
	} else {
		response.Status = 0
		response.Message = "Unknown Method"
	}
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	var arr_user []userStruct
	var response Response

	if r.Method == "DELETE" {
		var isId = r.FormValue("id")
		if isId != "" {
			_, err := db.Exec("DELETE FROM users WHERE id = ?", isId)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("Successfully delete data")
			response.Status = 1
			response.Message = "Successfully delete data"
		} else {
			response.Status = 0
			response.Message = "user id should be defined"
		}
	} else {
		response.Status = 0
		response.Message = "Unknown Method"
	}

	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/users", getUser)
	http.HandleFunc("/user_create", createUser)
	http.HandleFunc("/user_update", editUser)
	http.HandleFunc("/user_delete", deleteUser)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var response Response
		var arr_user []userStruct
		response.Status = 1
		response.Message = "Welcome to web api"
		response.Data = arr_user
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Running on 127.0.0.1:8181")
	http.ListenAndServe(":8181", nil)
}
