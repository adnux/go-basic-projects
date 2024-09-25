package bank

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/adnux/go-basic-projects/fileOps"
)

const accountBalanceFileName = "balance.txt"
const balanceMessage = "Your balance is %.2f"

func StartBankApp() {
	var accountBalance, error = fileOps.ReadFloatFromFile(accountBalanceFileName)
	if error != nil {
		fmt.Printf(Red+"ERROR: %v "+Reset+"\n", error)
		// panic(error)
		return
	}

	fmt.Println(Blue + "Welcome to the Go Bank App" + Reset)
	fmt.Println("Reach us 24/7 on", randomdata.PhoneNumber())

	for choice := 0; choice != 4; choice = 0 {

		PresentBankOptions()

		fmt.Print(Cyan + "Enter your choice: " + Reset)
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf(Yellow+balanceMessage+Reset+"\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Printf(Magenta + "Enter the amount you want to deposit: " + Reset)
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fileOps.WriteFloatToFile(accountBalanceFileName, accountBalance)
			fmt.Printf("You have deposited %.2f\n", depositAmount)
			fmt.Printf(Yellow+balanceMessage+Reset+"\n", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Println("Enter the amount you want to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println(Yellow + "Invalid withdraw amount. Must be greater than 0" + Reset)
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println(Red + "Insufficient funds" + Reset)
				continue
			}
			accountBalance -= withdrawAmount
			fileOps.WriteFloatToFile(accountBalanceFileName, accountBalance)
			fmt.Printf("You have withdrawn %.2f\n", withdrawAmount)
			fmt.Printf(Yellow+balanceMessage+Reset+"\n", accountBalance)
		case 4:
			fmt.Println("Thank you for using the Go Bank App")
			return
		default:
			fmt.Println("Invalid choice")
		}

	}
}
