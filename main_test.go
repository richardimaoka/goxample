package main

import (
	"fmt"
	"testing"
)

func validateAllocatedBudgetsTotal(budget int, allocatedBudgets []int, t *testing.T) {
	allocatedTotal := CalculateTotalBudget(allocatedBudgets)
	if budget != allocatedTotal {
		t.Errorf("budget = %d is different from total of allocated budgets = %d, where allocated budgets are = %v", budget, allocatedTotal, allocatedBudgets)
	}
}

func validateMinimumAllocation(budget int, allocatedBudgets []int, t *testing.T) {
	for _, v := range allocatedBudgets {
		allocation := float64(v) / float64(budget)
		if allocation < 0.1 {
			t.Errorf("allocation = %v is smaller than 10 percent", allocation)
		}
	}
}

func TestAll(t *testing.T) {
	testInputs := []Input{
		{Budget: 10 * MAN, CPAs: []float64{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}},
		{Budget: 100 * MAN, CPAs: []float64{300, 2402, 2500}},
	}

	for i, input := range testInputs {
		t.Run(fmt.Sprintf("test case: %d", i), func(t *testing.T) {
			allocatedBudgets := AllocationAlgorithm(input)
			validateAllocatedBudgetsTotal(input.Budget, allocatedBudgets, t)
			validateMinimumAllocation(input.Budget, allocatedBudgets, t)
		})
	}
}
