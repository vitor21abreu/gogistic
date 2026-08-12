package freight

// CalculateDistanceValue calcula o custo da distância.
func CalculateDistanceValue(distance, pricePerKm float64) float64 {
	return distance * pricePerKm
}
