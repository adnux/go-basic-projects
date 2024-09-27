package functions

import (
	"fmt"
)

type transformFn func(int) int

func StartFunctions() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("numbers =>", numbers)
	doubled := transformNumbers(numbers, double)
	fmt.Println("doubled =>", doubled)
	tripled := transformNumbers(numbers, triple)
	fmt.Println("tripled =>", tripled)
}

func transformNumbers(numbers []int, transform transformFn) []int {
	doubled := []int{}
	for _, number := range numbers {
		doubled = append(doubled, transform(number))
	}
	return doubled
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
