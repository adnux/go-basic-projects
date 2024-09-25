package main

import (
	"fmt"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	fmt.Println("Choose the app you want to run:")
	fmt.Println("1. Investment Calculator")
	fmt.Println("2. Profit Calculator")
	fmt.Println("3. Bank App")
	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		startInvestmentCalculator()
	case 2:
		startProfitCalculator()
	case 3:
		startBankApp()
	}

}
