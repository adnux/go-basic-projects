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

	doubled := TransformNumbers(&numbers, Double)
	fmt.Println("doubled =>", doubled)

	tripled := TransformNumbers(&numbers, Triple)
	fmt.Println("tripled =>", tripled)

	transformedFnOne := GetTransformerFunction(&numbers)
	transformedNumbersOne := TransformNumbers(&numbers, transformedFnOne)
	fmt.Println("transformedNumbersOne =>", transformedNumbersOne)

	transformedFnTwo := GetTransformerFunction(&moreNumbers)
	transformedTwo := TransformNumbers(&moreNumbers, transformedFnTwo)
	fmt.Println("transformedTwo =>", transformedTwo)

	anonymousNumbers := TransformNumbers(&numbers, CreateTransformerFn(10))
	fmt.Println("anonymousNumbers =>", numbers, "to", anonymousNumbers)

	factorialNumber := Factorial(20)
	p := message.NewPrinter(language.Portuguese)
	p.Printf("factorialNumber(%d) => %d \n", 20, factorialNumber)

	slice := []int{1, 2, 3, 4, 5}
	sum := SumUp(slice...)
	fmt.Println("sum =>", sum)
}

func SumUp(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func TransformNumbers(numbers *[]int, transform TransformFn) []int {
	doubled := []int{}
	for _, number := range *numbers {
		doubled = append(doubled, transform(number))
	}
	return doubled
}

func GetTransformerFunction(numbers *[]int) TransformFn {
	if len(*numbers) == 0 {
		return nil
	}
	if (*numbers)[0] == 1 {
		return Double
	}
	return Triple
}

func CreateTransformerFn(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func Factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * Factorial(number-1)
}

func Double(number int) int {
	return number * 2
}

func Triple(number int) int {
	return number * 3
}
