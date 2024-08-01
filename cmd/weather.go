package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
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

func Forecast(loc string) {
	// sets default value to Albany, CA
	q := loc

	if len(os.Args) > 2 {
		// set 'q' location to user-specific location
		q = os.Args[2]
	}

	// INSERT API KEY HERE
	key := ""
	url := "http://api.weatherapi.com/v1/forecast.json?key=" + key + "&q=" + q + "&days=1&aqi=no&alerts=no"

	res, err := http.Get(url)

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

		// edit this to change rain threshold, must be a float64
		rain_threshold := 60.0
		if hour.ChanceOfRain < rain_threshold {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}

var forecastCmd = &cobra.Command{
	Use: "forecast",
	Short: "Forecast hourly weather.",
	Long: "Will show rainy days in red.",
	Run: func(cmd *cobra.Command, args []string) {
		location := "94706"
		if len(args) >= 2 {
			location = args[1]
		}

		Forecast(location)
	},
}

func init() {
	rootCmd.AddCommand(forecastCmd)
}