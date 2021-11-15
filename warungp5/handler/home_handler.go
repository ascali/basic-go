package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var baseURL = "http://localhost:1323"

func HomeHandler(c echo.Context) error {
	// Please note the the second parameter "home.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	var datax, err = ambil_data()
	if err != nil {
		// log.Print(err.Error())
	}

	var dataj, err2 = ambil_data_jenis()
	if err2 != nil {
		// log.Print(err.Error())
	}

	var datap, err3 = ambil_data_populer()
	if err3 != nil {
		// log.Print(err.Error())
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":         "HOME",
		"msg":          "Echo",
		"data":         datax,
		"data_jenis":   dataj,
		"data_populer": datap,
	})
}

func ambil_data() ([]menu, error) {
	var err error
	var client = &http.Client{}
	var data []menu

	request, err := http.NewRequest("GET", baseURL+"/baca_menu", nil)
	if err != nil {
		log.Print(err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		log.Print(err.Error())
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Print(err.Error())
	}

	return data, nil
}

func ambil_data_jenis() ([]jenis_struct, error) {
	var err error
	var client = &http.Client{}
	var data_jenis []jenis_struct

	request, err := http.NewRequest("GET", baseURL+"/baca_jenis_menu", nil)
	if err != nil {
		log.Print(err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		log.Print(err.Error())
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data_jenis)
	if err != nil {
		log.Print(err.Error())
	}

	return data_jenis, nil
}

func ambil_data_populer() ([]menu, error) {
	var err error
	var client = &http.Client{}
	var data []menu

	request, err := http.NewRequest("GET", baseURL+"/baca_menu_populer", nil)
	if err != nil {
		log.Print(err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		log.Print(err.Error())
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Print(err.Error())
	}

	return data, nil
}
