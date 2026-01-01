package mult

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

var WeatherMap = map[int]WeatherCondition{1: Clear, 2: Rain, 3: HeavyRain, 4: Snow}

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed float64
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	res := 1.0
	switch weather.Condition {
	case HeavyRain:
		res += 0.2
	case Snow:
		res += 0.15
	case Rain:
		res += 0.125
	}
	if weather.WindSpeed > 15 {
		res += 0.1
	}
	return res
}
