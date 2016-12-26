package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type BaseWeather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID 			int `json:"id"`
		Main 		string `json:"main"`
		Description string `json:"description"`
		Icon 		string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp 		float64 `json:"temp"`
		Pressure 	int `json:"pressure"`
		Humidity 	int `json:"humidity"`
		TempMin 	float64 `json:"temp_min"`
		TempMax 	float64 `json:"temp_max"`
	} `json:"main"`
	Name string `json:"name"`
}

func main() {
	API_KEY := "API_KEY"

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter zipcode: ")
	scanner.Scan()
	zipcode := scanner.Text()

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s&appid=%s&units=imperial", zipcode, API_KEY)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var record BaseWeather

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Temperature: ", record.Main.Temp)
}