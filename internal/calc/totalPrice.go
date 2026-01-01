package calc

import (
	"math"
	"time"

	. "github.com/middelmatigheid/taxi-price/internal/mult"
)

const (
	minPrice = 99.0
	maxPrice = 20000.0
)

func ApplyPriceLimits(price float64) float64 {
	if price < minPrice {
		return minPrice
	} else if maxPrice < price {
		return maxPrice
	} else {
		return price
	}
}

type PriceCalculator struct {
	TrafficClient TrafficClient
}

func (c *PriceCalculator) CalculatePrice(trip TripParameters, now time.Time, weather WeatherData, lat, lng float64) float64 {
	base := CalculateBasePrice(trip)
	timeMult := GetTimeMultiplier(now)
	weatherMult := GetWeatherMultiplier(weather)
	trafficMult := GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMult * weatherMult * trafficMult

	return ApplyPriceLimits(math.Round(finalPrice))
}
