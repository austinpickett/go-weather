package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type BaseWeather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Name string `json:"name"`
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	API_KEY := viper.GetString("API_KEY")
	API_URL := viper.GetString("API_URL")

	var inTE, outTE *walk.TextEdit
	MainWindow{
		Title:   "Weather App",
		MinSize: Size{400, 200},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "Check Weather",
				OnClicked: func() {
					zipcode := inTE.Text()

					url := fmt.Sprintf("%s?zip=%s&appid=%s&units=imperial", API_URL, zipcode, API_KEY)

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

					output := fmt.Sprintf("Temperature: %f", record.Main.Temp)

					outTE.SetText(output)
				},
			},
		},
	}.Run()
}
