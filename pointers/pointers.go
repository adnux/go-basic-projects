package pointers

import "fmt"

func StartAdulthood() {
	// scan the age from the user
	var age int
	agePointer := &age

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Age:", *agePointer)

	adultYears := getAdulthoodYears(*agePointer)
	fmt.Println("Years as an adult:", adultYears)
}

func getAdulthoodYears(age int) int {
	// *age = *age - 18
	return age - 18
}
