package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Current struct {
		TempF float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		}`json:"condition"`
	} `json:"current"`

	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64 `json:"time_epoch"`
				TempF float64 `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	// sets default value to Albany, CA
	q := "94706"

	if len(os.Args) >= 2 {
		// set 'q' location to user-specific location
		q = os.Args[1]
	}

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=99cdb051202a41f0b9b221441243107&q=" + q + "&days=1&aqi=no&alerts=no")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close();

	if res.StatusCode != 200 {
		panic("Error fetching data from Weather API.")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Map JSON response to weather struct
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	//fmt.Println(weather)

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	fmt.Printf("%s, %s: %.0fF, %s\n", location.Name, location.Country, current.TempF, current.Condition.Text)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		// only display future hours
		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf("%s - %.0fF, %.0f%%, %s\n", date.Format("03:04"), hour.TempF, hour.ChanceOfRain, hour.Condition.Text)
		
		// highlight hours with greater than 60% chance of rain
		if hour.ChanceOfRain < 60 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}