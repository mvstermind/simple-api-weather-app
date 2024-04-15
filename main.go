package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"weather-app/api"
	request "weather-app/handleRequest"
)

var url string = "https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s"

func main() {

	fmt.Printf("Type a city to get info about weather: ")
	b := bufio.NewReader(os.Stdin)

	cityName, err := b.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	cityName = strings.TrimSpace(cityName)
	frmtdUrl := fmt.Sprintf(url, cityName, api.ApiKey)

	body := request.ApiRequest(frmtdUrl)
	fmt.Print(body)
}
