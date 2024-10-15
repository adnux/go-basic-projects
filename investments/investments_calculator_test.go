package investments

import (
	"math"
	"testing"
)

func TestCalculator(t *testing.T) {
	t.Run("CalculateValue", CalculateValueTest)
	t.Run("CalculateFutureValue", CalculateFutureValueTest)

	t.Run("CalculateFutureAdjustedValue", CalculateFutureAdjustedValueTest)
}

func CalculateValueTest(t *testing.T) {
	// Example test case
	inflationRate := 2.5
	expectedReturnRate := 5.5
	investmentAmount := 10000.0
	years := 10

	expectedFutureValue := 17081.444584
	expectedFutureRealValue := 13343.997208

	futureValue, futureRealValue := CalculateValue(investmentAmount, expectedReturnRate, float64(years), inflationRate)
	if math.Abs(futureValue-expectedFutureValue) > 1e-6 || math.Abs(futureRealValue-expectedFutureRealValue) > 1e-6 {
		t.Errorf("CalculateValue(%f, %f, %d, %f) = %f, %f; want %f, %f", investmentAmount, expectedReturnRate, years, inflationRate, futureValue, futureRealValue, expectedFutureValue, expectedFutureRealValue)
	}
}

func CalculateFutureValueTest(t *testing.T) {
	investmentAmount := 1000.0
	expectedReturnRate := 5.0
	years := 10
	expectedFutureValue := 1628.894627

	futureValue := CalculateFutureValue(investmentAmount, expectedReturnRate, float64(years))
	if math.Abs(futureValue-expectedFutureValue) > 1e-6 {
		t.Errorf("CalculateFutureValue(%f, %f, %d) = %f; want %f", investmentAmount, expectedReturnRate, years, futureValue, expectedFutureValue)
	}
}

func CalculateFutureAdjustedValueTest(t *testing.T) {
	futureValue := 1628.89
	inflationRate := 2.5
	years := float64(10)
	expectedFutureRealValue := 1272.486265

	futureRealValue := CalculateFutureAdjustedValue(futureValue, inflationRate, years)
	if math.Abs(futureRealValue-expectedFutureRealValue) > 1e-6 {
		t.Errorf("CalculateFutureAdjustedValue(%f, %f, %f) = %f; want %f", futureValue, inflationRate, years, futureRealValue, expectedFutureRealValue)
	}
}
