package main

import (
	"testing"
)

func validateAllocatedBudgetsTotal(budget int, allocatedBudgets []int, t *testing.T) {
	allocatedTotal := CalculateTotalBudget(allocatedBudgets)
	if budget != allocatedTotal {
		t.Errorf("budget = %d is different from total of allocated budgets = %d, where allocated budgets are = %v", budget, allocatedTotal, allocatedBudgets)
	}
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

func TestAll(t *testing.T) {
	testInputs := []Input{
		{1000000, []float64{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}},
	}

	for _, input := range testInputs {
		allocatedBudgets := AllocationAlgorithm(input)
		validateAllocatedBudgetsTotal(input.Budget, allocatedBudgets, t)
	}
}
