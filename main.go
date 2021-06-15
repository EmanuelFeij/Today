package main

import (
	"fmt"

	"github.com/EmanuelFeij/Today/weather"
)

func main() {
	prevision, err := weather.GetWeather("CoImBra", true)
	if err != nil {
		return
	}
	fmt.Println(prevision)
}
