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

func validateAllocatedBudgetsTotal(budget int, allocatedBudgets []int) bool {
	allocatedTotal := CalculateTotal(allocatedBudgets)
	return budget == allocatedTotal
}

func validateMinimumAllocation(budget int, allocatedBudgets []int) bool {
	for _, v := range allocatedBudgets {
		proportion := float64(v) / float64(budget)
		if proportion < 0.1 {
			return false
		}
	}
	return true
}
