package main

import "fmt"

func CalculateTotal(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

func CalculateInverse(slice []float64) []float64 {
	inversed := make([]float64, len(slice))
	for i, v := range slice {
		inversed[i] = 1 / v
	}
	return inversed
}

func main() {

	slice := []float64{10, 10, 10}

	fmt.Println(Calculat(slice))
}
