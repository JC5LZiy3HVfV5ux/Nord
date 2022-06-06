package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JC5LZiy3HVfV5ux/nord/pkg/openweather"
)

func main() {
	client, err := openweather.NewClient("a8ca8ed93c0665175e6a269dbf241f8c")
	if err != nil {
		log.Fatal(err)
	}

	if err := client.SetLang("ru"); err != nil {
		log.Println(err)
	}

	if err := client.SetUnit("metric"); err != nil {
		log.Println(err)
	}

	data := &openweather.CurrentWeatherData{}

	if err := client.CurrentWeather().CurrentByCityId(context.Background(), data, 2172797); err != nil {
		log.Fatal(err)
	}

	fmt.Println("---------------------------------------------------\\")

	fmt.Printf("lon: %f\n", data.Coord.Lon)

	fmt.Printf("lat: %f\n", data.Coord.Lat)

	fmt.Printf("weatherId: %d\n", data.Weather[0].Id)

	fmt.Printf("main: %s\n", data.Weather[0].Main)

	fmt.Printf("description: %s\n", data.Weather[0].Description)

	fmt.Printf("icon: %s\n", data.Weather[0].Icon)

	fmt.Printf("icon url: %s\n", data.Weather[0].UrlIconWeather())

	fmt.Printf("temp: %f\n", data.Main.Temp)

	fmt.Printf("pressure: %f\n", data.Main.Pressure)

	fmt.Printf("humidity: %d\n", data.Main.Humidity)

	fmt.Printf("tempMin: %f\n", data.Main.TempMin)

	fmt.Printf("tempMax: %f\n", data.Main.TempMax)

	fmt.Printf("windSpeed: %f\n", data.Wind.Speed)

	fmt.Printf("windDeg: %f\n", data.Wind.Deg)

	fmt.Printf("clouds: %d\n", data.Clouds.All)

	fmt.Printf("country: %s\n", data.Sys.Country)

	fmt.Printf("cityName: %s\n", data.Name)

	fmt.Printf("---------------------------------------------------/\n")
}
