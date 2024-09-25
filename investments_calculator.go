package main

import (
	"fmt"
	"math"
	"strings"
)

func startInvestmentCalculator() {

	fmt.Print(("Enter the inflation rate: (2.5 default)"))
	var inflationRate float64 = 2.5
	fmt.Scan(&inflationRate)
	fmt.Print(("Enter the expected return rate: (5.5 default)"))
	var expectedReturnRate float64 = 5.5
	fmt.Scan(&expectedReturnRate)
	investmentCalculator(inflationRate, expectedReturnRate)
}

func investmentCalculator(inflationRate float64, expectedReturnRate float64) {
	fmt.Print("Inflation Rate: ", inflationRate, "\n")
	fmt.Print("Expected Return Rate: ", expectedReturnRate, "\n")
	var investmentAmount float64
	var years float64

	outputText("Expected Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateValue(investmentAmount, expectedReturnRate, years, inflationRate)
	// futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := calculateFutureAdjustedValue(futureValue, inflationRate, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFutureValue := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFutureAdjustedValue := fmt.Sprintf("Adjusted Future Value: %.2f\n", futureRealValue)
	outputText(formattedFutureValue, formattedFutureAdjustedValue)
}

func calculateValue(investmentAmount float64, expectedReturnRate float64, years float64, inflationRate float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, years float64) float64 {
	return investmentAmount * math.Pow(1+expectedReturnRate/100, years)
}

func calculateFutureAdjustedValue(futureValue float64, inflationRate float64, years float64) float64 {
	return futureValue / math.Pow(1+inflationRate/100, years)
}

func outputText(text ...string) {
	fmt.Print(strings.Join(text, ""))
}
