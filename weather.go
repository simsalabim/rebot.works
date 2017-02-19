package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type WeatherApiResponse struct {
	Main MainResponse `json:"main"`
}

type MainResponse struct {
	Temp float64 `json:"temp"`
}

func weatherJSON() string {
	data_chi := weatherData("Chicago,IL")
	data_sea := weatherData("Seattle,WA")
	data_pa := weatherData("Palo Alto,CA")

	response := fmt.Sprintf(`
	{
		"%s": { "celcius": %f, "fahrenheit": %f },
		"%s": { "celcius": %f, "fahrenheit": %f },
		"%s": { "celcius": %f, "fahrenheit": %f }
	}
	`,
		"Chicago", celcius(data_chi.Main.Temp), fahrenheit(data_chi.Main.Temp),
		"Seattle", celcius(data_sea.Main.Temp), fahrenheit(data_sea.Main.Temp),
		"Palo Alto", celcius(data_pa.Main.Temp), fahrenheit(data_pa.Main.Temp),
	)

	return response
}

func weatherData(location string) *WeatherApiResponse {
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + url.QueryEscape(location) + "&appid=" + os.Getenv("WEATHER_APP_ID")
	resp, err := http.Get(url)
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	var api_response = new(WeatherApiResponse)
	json.Unmarshal(body, &api_response)
	return api_response
}

func celcius(temp float64) float64 {
	return temp - 273.15
}

func fahrenheit(temp float64) float64 {
	return 9/5*(temp-273.15) + 32
}
