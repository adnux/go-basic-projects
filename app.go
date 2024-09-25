package main

import (
	"fmt"

	"github.com/adnux/go-basic-projects/bank"
	"github.com/adnux/go-basic-projects/investments"
	"github.com/adnux/go-basic-projects/pointers"
	"github.com/adnux/go-basic-projects/profits"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	fmt.Println("Choose the app you want to run:")
	fmt.Println("1. Investment Calculator")
	fmt.Println("2. Profit Calculator")
	fmt.Println("3. Bank App")
	fmt.Println("4. Pointers")
	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		investments.StartInvestmentCalculator()
	case 2:
		profits.StartProfitCalculator()
	case 3:
		bank.StartBankApp()
	case 4:
		pointers.StartAdulthood()
	}

}
