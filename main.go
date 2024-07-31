package main

import (
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Country struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"country"`
}

func main() {
	res, err := http.Get("")

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

	fmt.Println(string(body))
}