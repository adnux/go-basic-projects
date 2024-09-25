package main

import "fmt"

func PresentOptions() {
	fmt.Println(Green + "What do you want to do?" + Reset)
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}
