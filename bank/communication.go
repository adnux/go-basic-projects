package bank

import (
	"fmt"

	"github.com/adnux/go-basic-projects/colors"
)

func PresentBankOptions() {
	fmt.Println(colors.Green + "What do you want to do?" + colors.Reset)
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}
