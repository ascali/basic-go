package handler

import (
	"fmt"
	"log"
	_ "mysql-master"
	"net/http"
	"warung/server"

	"github.com/labstack/echo"
)

type Menu struct {
	Id          string `json:"id"`
	TypeId      string `json:"type_id"`
	TypeName    string `json:"type_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       string `json:"price"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Menu `json:"data"`
}

type TypeMenuMaster struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ResponseType struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    []TypeMenuMaster `json:"data"`
}

func GetDataTypeMenu(c echo.Context) error {
	var EachTypeMenu TypeMenuMaster
	var DataTypeMenu []TypeMenuMaster
	var responseTypeMenuMaster ResponseType

	db := server.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM warung_lengko.type_master ORDER BY name ASC")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {

		var err = rows.Scan(&EachTypeMenu.Id, &EachTypeMenu.Name)

		if err != nil {
			log.Fatal(err.Error())
		} else {
			DataTypeMenu = append(DataTypeMenu, EachTypeMenu)
			fmt.Println(DataTypeMenu)
		}
	}
	responseTypeMenuMaster.Status = 1
	responseTypeMenuMaster.Message = "Success"
	responseTypeMenuMaster.Data = DataTypeMenu

	return c.JSON(http.StatusOK, responseTypeMenuMaster)
}

func GetDataMenu(c echo.Context) error {
	var each Menu
	var DataMenu []Menu
	var response Response

	db := server.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT menus.id, type_master.id as type_id, type_master.name as type_name, menus.name, menus.description, menus.image, menus.price FROM warung_lengko.menus menus, warung_lengko.type_master type_master WHERE menus.type_id = type_master.id")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		var err = rows.Scan(&each.Id, &each.TypeId, &each.TypeName, &each.Name, &each.Description, &each.Image, &each.Price)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			DataMenu = append(DataMenu, each)
			fmt.Println(DataMenu)
		}
	}
	response.Status = 1
	response.Message = "Success"
	response.Data = DataMenu

	return c.JSON(http.StatusOK, response)
}
