package main

import (
	"fmt"
	"strconv"
	"time"

	. "github.com/middelmatigheid/taxi-price/internal/calc"
	. "github.com/middelmatigheid/taxi-price/internal/mult"
)

func main() {
	calculator := PriceCalculator{
		TrafficClient: &RealTrafficClient{}, // A real traffic client is being used in a production version but we will replace it with a dummy
	}

	fmt.Println("Hello, it is a simplify version of a taxi price calculator. Here you can see how the price varies due to different conditions")

	var input string
	var num float64
	var err error
	var trip TripParameters
	var weather WeatherData
	var currentTime time.Time
	var lat, lon float64

	for true {
		fmt.Println()

		// Getting ride distance
		fmt.Print("Ride distance in kilometers: ")
		fmt.Scan(&input)
		if input == "exit" {
			break
		}
		num, err = strconv.ParseFloat(input, 64)
		if err != nil || num <= 0 || num > 10000 {
			fmt.Println("Incorrect value, please try again")
			continue
		}
		trip.Distance = num

		// Getting ride duration
		fmt.Print("Ride duration in minutes: ")
		fmt.Scan(&input)
		if input == "exit" {
			break
		}
		num, err = strconv.ParseFloat(input, 64)
		if err != nil || num <= 0 || num > 10000 {
			fmt.Println("Incorrect value, please try again")
			continue
		}
		trip.Duration = num

		// Getting current time
		fmt.Print("Current time in a 15:04 format: ")
		fmt.Scan(&input)
		if input == "exit" {
			break
		}
		currentTime, err = time.Parse("15:04", input)
		if err != nil {
			fmt.Println("Incorrect value, please try again")
			continue
		}

		// Getting weather condition
		fmt.Print("Weather condition (1 - Clear, 2 - Rain, 3 - Heavy rain, 4 - Snow): ")
		fmt.Scan(&input)
		if input == "exit" {
			break
		}
		num, err = strconv.ParseFloat(input, 64)
		if err != nil || num < 1 || num > 4 {
			fmt.Println("Incorrect value, please try again")
			continue
		}
		weather.Condition = WeatherMap[int(num)]

		// Getting wind speed
		fmt.Print("Wind speed: ")
		fmt.Scan(&input)
		if input == "exit" {
			break
		}
		num, err = strconv.ParseFloat(input, 64)
		if err != nil || num < 0 || num > 100 {
			fmt.Println("Incorrect value, please try again")
			continue
		}
		weather.WindSpeed = num

		// Getting location
		fmt.Print("Specify your location in \"lat lot\" format: ")
		// Getting latitude
		fmt.Scan(&input)
		if input == "exit" {
			break
		}
		num, err = strconv.ParseFloat(input, 64)
		if err != nil || num < -90 || num > 90 {
			fmt.Println("Incorrect value, please try again")
			continue
		}
		lat = num
		// Getting longtitude
		fmt.Scan(&input)
		if input == "exit" {
			break
		}
		num, err = strconv.ParseFloat(input, 64)
		if err != nil || num < -180 || num > 180 {
			fmt.Println("Incorrect value, please try again")
			continue
		}
		lon = num

		// Calculating total price
		price := calculator.CalculatePrice(trip, currentTime, weather, lat, lon)
		fmt.Printf("Total price: %.0f руб.\n", price)
	}
}
