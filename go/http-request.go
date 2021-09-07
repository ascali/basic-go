package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var BaseUrl = "http://localhost:8080"

type FoodMenu struct {
	Id       string
	MenuName string
	Price    int
}

func GetApi(menu string) (FoodMenu, error) {
	var err error
	var client = &http.Client{}
	var data FoodMenu
	var param = url.Values{}
	param.Set("MenuName", menu)

	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("GET", BaseUrl+"/search_menu", payload)
	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "Application/json")
	response, err := client.Do(request)
	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	var menu, err = GetApi("Rumbah%20Kucur")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Menu", menu)
}
