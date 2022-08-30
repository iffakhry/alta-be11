package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Peoples struct {
	Count  int            `json:"count"`
	Result []SwapiPeoples `json:"results"`
}

type SwapiPeoples struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	EyeColor  string   `json:"eye_color"`
	SkinColor string   `json:"skin_color"`
	Films     []string `json:"films"`
}

func main() {
	response, err := http.Get("https://swapi.dev/api/people")
	if err != nil {
		log.Fatal("error get")
	}

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var people9 Peoples
	json.Unmarshal(responseData, &people9) // merubah json to object
	// json.Marshal() //merubah object to json

	for _, v := range people9.Result {
		fmt.Println("name", v.Name)
	}

	fmt.Println("Name", people9.Result[1].Name)
	fmt.Println("Height", people9.Result[1].Height)
	fmt.Println("Skin Color", people9.Result[1].SkinColor)
	fmt.Println("Skin Color", people9.Result[1].EyeColor)

	for _, v := range people9.Result[1].Films {
		fmt.Println("film", v)
	}
}
