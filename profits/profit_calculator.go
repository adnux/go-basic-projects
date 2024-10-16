package profits

import (
	"errors"
	"fmt"
	"os"
)

func StartProfitCalculator() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := CalculateFinancialValues(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("EBT to Profit Ratio: %.3f\n", ratio)

	storeResultsInFile(ebt, profit, ratio)

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput < 0 {
		return 0, errors.New("value must be a positive number")
	}
	return userInput, nil

}
func storeResultsInFile(ebt float64, profit float64, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	fmt.Println("Storing results in file...")
	os.WriteFile("profits.txt", []byte(results), 0644)
}

func CalculateFinancialValues(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func CalculateFinancialValuesComposed(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = CalculateEbt(revenue, expenses)
	profit = CalculateProfit(ebt, taxRate)
	ratio = CalculateRatio(ebt, profit)
	return ebt, profit, ratio
}

func CalculateEbt(revenue float64, expenses float64) float64 {
	ebt := revenue - expenses
	return ebt
}

func CalculateProfit(ebt float64, taxRate float64) float64 {
	profit := ebt * (1 - taxRate/100)
	return profit
}

func CalculateRatio(ebt float64, profit float64) float64 {
	ratio := ebt / profit
	return ratio
}
