package bank

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/adnux/go-basic-projects/colors"
	"github.com/adnux/go-basic-projects/fileOps"
)

const accountBalanceFileName = "balance.txt"
const balanceMessage = "Your balance is %.2f"

func StartBankApp() {
	var accountBalance, error = fileOps.ReadFloatFromFile(accountBalanceFileName, true)
	if error != nil {
		fmt.Printf(colors.Red+"ERROR: %v "+colors.Reset+"\n", error)
		// panic(error)
		return
	}

	fmt.Println(colors.Blue + "Welcome to the Go Bank App" + colors.Reset)
	fmt.Println("Reach us 24/7 on", randomdata.PhoneNumber())

	for choice := 0; choice != 4; choice = 0 {

		PresentBankOptions()

		fmt.Print(colors.Cyan + "Enter your choice: " + colors.Reset)
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf(colors.Yellow+balanceMessage+colors.Reset+"\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Printf(colors.Magenta + "Enter the amount you want to deposit: " + colors.Reset)
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fileOps.WriteFloatToFile(accountBalanceFileName, accountBalance)
			fmt.Printf("You have deposited %.2f\n", depositAmount)
			fmt.Printf(colors.Yellow+balanceMessage+colors.Reset+"\n", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Println("Enter the amount you want to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println(colors.Yellow + "Invalid withdraw amount. Must be greater than 0" + colors.Reset)
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println(colors.Red + "Insufficient funds" + colors.Reset)
				continue
			}
			accountBalance -= withdrawAmount
			fileOps.WriteFloatToFile(accountBalanceFileName, accountBalance)
			fmt.Printf("You have withdrawn %.2f\n", withdrawAmount)
			fmt.Printf(colors.Yellow+balanceMessage+colors.Reset+"\n", accountBalance)
		case 4:
			fmt.Println("Thank you for using the Go Bank App")
			return
		default:
			fmt.Println("Invalid choice")
		}

	}
}
