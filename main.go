package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"weather-app/api"
)

var url string = "https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s"

func main() {

	fmt.Printf("Type a city to get info about weather: ")
	b := bufio.NewReader(os.Stdin)

	cityName, err := b.ReadString('\n')
	cityName = strings.TrimSpace(cityName)
	frmtdUrl := fmt.Sprintf(url, cityName, api.ApiKey)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	res, err := http.Get(frmtdUrl)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Print(string(body))
}
