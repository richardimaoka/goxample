package main

func CalculateTotal(slice []float64) float64 {
	total := 0.0
	for _, v := range slice {
		total += v
	}
	return total
}
