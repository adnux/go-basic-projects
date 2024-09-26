package main

import (
	"fmt"

	"github.com/adnux/go-basic-projects/bank"
	"github.com/adnux/go-basic-projects/investments"
	"github.com/adnux/go-basic-projects/notes"
	"github.com/adnux/go-basic-projects/pointers"
	"github.com/adnux/go-basic-projects/profits"
	"github.com/adnux/go-basic-projects/structs"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	fmt.Println("Choose the app you want to run:")
	fmt.Println("1. Investment Calculator")
	fmt.Println("2. Profit Calculator")
	fmt.Println("3. Bank App")
	fmt.Println("4. Pointers")
	fmt.Println("5. Structs")
	fmt.Println("6. Notes")
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
	case 5:
		structs.StartUserStructs()
	case 6:
		notes.StartNotes()
	}

}
