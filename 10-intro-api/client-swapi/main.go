package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SwapiPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	EyeColor  string   `json:"eye_color"`
	SkinColor string   `json:"skin_color"`
	Films     []string `json:"films"`
}

func main() {
	response, err := http.Get("https://swapi.dev/api/people/1")
	if err != nil {
		log.Fatal("error get")
	}

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var people9 SwapiPeople
	json.Unmarshal(responseData, &people9) // merubah json to object
	// json.Marshal() //merubah object to json

	fmt.Println("Name", people9.Name)
	fmt.Println("Height", people9.Height)
	fmt.Println("Skin Color", people9.SkinColor)
	fmt.Println("Skin Color", people9.EyeColor)

	for _, v := range people9.Films {
		fmt.Println("film", v)
	}
}
