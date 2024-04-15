package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"weather-app/api"
	request "weather-app/handleRequest"
	WeatherJson "weather-app/json"
)

var url string = "https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s"
var Weather WeatherJson.S

func main() {

	fmt.Printf("Type a city to get info about weather: ")
	b := bufio.NewReader(os.Stdin)

	cityName, err := b.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	cityName = strings.TrimSpace(cityName)
	frmtdUrl := fmt.Sprintf(url, cityName, api.ApiKey) //combine user input with api key

	body := request.ApiRequest(frmtdUrl) // read api using combined link

	err = json.Unmarshal([]byte(body), &Weather) // read json
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}
	PrintJsonContent(Weather) // print content
}

// i have lieterall no clue how to export this
func PrintJsonContent(Weather WeatherJson.S) {

	fmt.Printf("Temperature: %v C\n", Weather.Main.Temp)
	fmt.Printf("Max Temperature: %v C\n", Weather.Main.TempMax)
	fmt.Printf("Min Temperature: %v C\n", Weather.Main.TempMin)
	fmt.Printf("Feels Like: %v C\n", Weather.Main.FeelsLike)
	fmt.Printf("Humidity: %v%%\n", Weather.Main.Humidity)
	fmt.Printf("Pressure: %v hPa\n", Weather.Main.Pressure)
	fmt.Printf("Wind Speed: %v Km\\H\n", Weather.Wind.Speed)
}
