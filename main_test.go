package main

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func validateAllocatedBudgetsTotal(budget float64, allocatedBudgets []float64) bool {
	allocatedTotal := CalculateTotal(allocatedBudgets)
	return math.Abs((budget-allocatedTotal)/budget) < 0.000001
}
