# 🌧️ Go Weather CLI 🌧️

Get weather forecasts in your terminal!

## Description

This is a CLI built in Go that fetches data from WeatherAPI and displays it in your terminal. Rainer days will be highlighted in red, or any color that you'd like. Location can be specified as well, but the program defaults to Albany, CA. 

## Screenshots

* ![Manhattan Forecast](https://ibb.co/sjvYYRC)

### Dependencies

* Go
* Free API key from https://www.weatherapi.com/

### Installing

* Clone the repository with `git clone https://github.com/davidl21/go_cli.git`
* Install all dependencies by typing `go mod tidy`
* Insert API key into weather.go

### Executing program

* In terminal and the go_cli directory, run `go run main.go forecast`.
* To specify a location, run `go run main.go forecast {location}`, replacing location with your specified location. 

## Customization

* If you want to change the rain percentage threshold for changing the text color, you can change the value of `rain_threshold` in weather.go.
* If you want to change the text highlight color for rainy days, you can use a different color function in weather.go from `github.com/fatih/color`. By default, the CLI uses `color.Red()` for rainy days. 

## Authors

Please contact me for any questions! 

davidl21@berkeley.edu
[linkedin.com/](https://www.linkedin.com/in/davidl21/)
