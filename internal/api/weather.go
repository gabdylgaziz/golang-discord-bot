package api

import (
	"discord/internal/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWeather(city string) (result entity.WeatherResponse) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=a1285e1455014769a06124313240102&q=%s&lang=ru", city)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var weather entity.Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	result = entity.WeatherResponse{
		Name:      weather.Location.Name,
		Region:    weather.Location.Region,
		Country:   weather.Location.Country,
		Localtime: weather.Location.Localtime,
		TempC:     weather.Current.TempC,
		TempF:     weather.Current.TempF,
		Text:      weather.Current.Condition.Text,
		WindKph:   weather.Current.WindKph,
		Humidity:  weather.Current.Humidity,
	}

	return
}
