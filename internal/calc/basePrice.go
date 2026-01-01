package calc

const (
	pricePerKm     = 10.0
	pricePerMinute = 2.0
)

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(t TripParameters) float64 {
	return t.Distance*pricePerKm + t.Duration*pricePerMinute
}
