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

	var data openweather.ListGeocodingData

	if err := client.Geocoding().GeocodingByCityName(context.Background(), &data, "London", 3); err != nil {
		log.Fatal(err)
	}

	fmt.Println("---------------------------------------------------\\")
	for k, v := range data {
		fmt.Printf("---\nGeocoding number: %d\n", k)

		fmt.Printf("name: %s\n", v.Name)

		fmt.Printf("lon: %f\n", v.Lon)

		fmt.Printf("lat: %f\n", v.Lat)

		fmt.Printf("country: %s\n", v.Country)
	}
	fmt.Printf("---------------------------------------------------/\n")

	zipData := &openweather.ZipGeocodingData{}

	if err := client.Geocoding().GeocodingByZip(context.Background(), zipData, "634000,ru"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Geocoding by zip code\n")

	fmt.Println("---------------------------------------------------\\")

	fmt.Printf("zip: %s\n", zipData.Zip)

	fmt.Printf("name: %s\n", zipData.Name)

	fmt.Printf("lon: %f\n", zipData.Lon)

	fmt.Printf("lat: %f\n", zipData.Lat)

	fmt.Printf("country: %s\n", zipData.Country)

	fmt.Printf("---------------------------------------------------/\n")
}
