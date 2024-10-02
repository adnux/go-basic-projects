package functions

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type TransformFn func(int) int

func StartFunctions() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("numbers =>", numbers)
	moreNumbers := []int{6, 7, 8, 9, 10}
	fmt.Println("moreNumbers =>", moreNumbers)

	doubled := transformNumbers(&numbers, double)
	fmt.Println("doubled =>", doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println("tripled =>", tripled)

	transformedFnOne := getTransformerFunction(&numbers)
	transformedNumbersOne := transformNumbers(&numbers, transformedFnOne)
	fmt.Println("transformedNumbersOne =>", transformedNumbersOne)

	transformedFnTwo := getTransformerFunction(&moreNumbers)
	transformedTwo := transformNumbers(&moreNumbers, transformedFnTwo)
	fmt.Println("transformedTwo =>", transformedTwo)

	anonymousNumbers := transformNumbers(&numbers, createTransformerFn(10))
	fmt.Println("anonymousNumbers =>", numbers, "to", anonymousNumbers)

	factorialNumber := factorial(20)
	p := message.NewPrinter(language.Portuguese)
	p.Printf("factorialNumber(%d) => %d \n", 20, factorialNumber)

	slice := []int{1, 2, 3, 4, 5}
	sum := sumUp(slice...)
	fmt.Println("sum =>", sum)
}

func sumUp(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func transformNumbers(numbers *[]int, transform TransformFn) []int {
	doubled := []int{}
	for _, number := range *numbers {
		doubled = append(doubled, transform(number))
	}
	return doubled
}

func getTransformerFunction(numbers *[]int) TransformFn {
	if len(*numbers) == 0 {
		return nil
	}
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func createTransformerFn(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
