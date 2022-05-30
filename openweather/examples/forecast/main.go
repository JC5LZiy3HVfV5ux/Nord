package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JC5LZiy3HVfV5ux/openweather-cache-server/openweather"
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

	data := &openweather.ForecastData{}

	// Tomsk
	if err := client.Forecast().ForecastByZip(context.Background(), data, "634000,ru", 3); err != nil {
		log.Fatal(err)
	}

	fmt.Println("---------------------------------------------------\\")

	fmt.Printf("lon: %f\n", data.City.Coord.Lon)

	fmt.Printf("lat: %f\n", data.City.Coord.Lat)

	for k, v := range data.List {
		fmt.Printf("---\nForecast number: %d\n", k)

		fmt.Printf("weatherId: %d\n", v.Weather[0].Id)

		fmt.Printf("main: %s\n", v.Weather[0].Main)

		fmt.Printf("description: %s\n", v.Weather[0].Description)

		fmt.Printf("icon: %s\n", v.Weather[0].Icon)

		fmt.Printf("icon url: %s\n", v.Weather[0].UrlIconWeather())

		fmt.Printf("temp: %f\n", v.Main.Temp)

		fmt.Printf("pressure: %f\n", v.Main.Pressure)

		fmt.Printf("humidity: %d\n", v.Main.Humidity)

		fmt.Printf("tempMin: %f\n", v.Main.TempMin)

		fmt.Printf("tempMax: %f\n", v.Main.TempMax)

		fmt.Printf("windSpeed: %f\n", v.Wind.Speed)

		fmt.Printf("windDeg: %f\n", v.Wind.Deg)

		fmt.Printf("clouds: %d\n", v.Clouds.All)
	}

	fmt.Printf("---\n")

	fmt.Printf("country: %s\n", data.City.Country)

	fmt.Printf("cityName: %s\n", data.City.Name)

	fmt.Printf("---------------------------------------------------/\n")
}
