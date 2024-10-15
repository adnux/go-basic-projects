package profits

import (
	"math"
	"testing"
)

// Example test case
var revenue = 1000.0
var expenses = 500.0
var taxRate = 25.0

// Expected results
var expectedEbt = 500.0
var expectedProfit = 375.0
var expectedRatio = 1.3333333333333333

func TestCalculateFinancialValues(t *testing.T) {

	ebt, profit, ratio := CalculateFinancialValues(revenue, expenses, taxRate)

	if math.Nextafter(ebt, expectedEbt) != expectedEbt ||
		math.Nextafter(profit, expectedProfit) != expectedProfit ||
		math.Nextafter(ratio, expectedRatio) != expectedRatio {
		t.Errorf("CalculateFinancialValues(%f, %f, %f) = %f, %f, %f; want %f, %f, %f", revenue, expenses, taxRate, ebt, profit, ratio, expectedEbt, expectedProfit, expectedRatio)
	}
}

func TestCalculateFinancialValuesComposed(t *testing.T) {

	ebt, profit, ratio := CalculateFinancialValuesComposed(revenue, expenses, taxRate)

	if math.Nextafter(ebt, expectedEbt) != expectedEbt ||
		math.Nextafter(profit, expectedProfit) != expectedProfit ||
		math.Nextafter(ratio, expectedRatio) != expectedRatio {
		t.Errorf("CalculateFinancialValuesComposed(%f, %f, %f) = %f, %f, %f; want %f, %f, %f", revenue, expenses, taxRate, ebt, profit, ratio, expectedEbt, expectedProfit, expectedRatio)
	}
}

func TestCalculateEbt(t *testing.T) {

	ebt := CalculateEbt(revenue, expenses)
	if math.Nextafter(ebt, expectedEbt) != expectedEbt {
		t.Errorf("CalculateEbt(%f, %f) = %f; want %f", revenue, expenses, ebt, expectedEbt)
	}
}

func TestCalculateProfit(t *testing.T) {
	ebt := CalculateEbt(revenue, expenses)

	profit := CalculateProfit(ebt, taxRate)
	if math.Nextafter(profit, expectedProfit) != expectedProfit {
		t.Errorf("CalculateProfit(%f, %f) = %f; want %f", ebt, taxRate, profit, expectedProfit)
	}
}

func TestCalculateRatio(t *testing.T) {
	ebt := CalculateEbt(revenue, expenses)
	profit := CalculateProfit(ebt, taxRate)

	ratio := CalculateRatio(ebt, profit)

	if math.Nextafter(ratio, expectedRatio) != expectedRatio {
		t.Errorf("CalculateRatio(%f, %f) = %f; want %f", ebt, profit, ratio, expectedRatio)
	}
}
