package main

import "fmt"

func CalculateTotalBudget(allocatedBudgets []int) int {
	total := 0
	for _, v := range allocatedBudgets {
		total += v
	}
	return total
}

func CalculateTotalProportion(proportions []float64) float64 {
	total := 0.0
	for _, v := range proportions {
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

type Input struct {
	Budget int
	CPAs   []float64
}

func AllocationAlgorithm(input Input) []int {
	inverse := CalculateInverse(input.CPAs)
	inverseTotal := CalculateTotalProportion(inverse)

	naiveAllocations := make([]float64, len(inverse))
	for i, v := range inverse {
		naiveAllocations[i] = v / inverseTotal
	}

	naiveAllocationBudgets := make([]int, len(naiveAllocations))
	for i, v := range naiveAllocations {
		naiveAllocationBudgets[i] = int(v * float64(input.Budget))
	}

	return naiveAllocationBudgets
}

func main() {
	input := Input{10000000, []float64{10, 10, 10}}

	fmt.Println(AllocationAlgorithm(input))
}
