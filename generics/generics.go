package generics

import "fmt"

func StartGenerics() {
	first, second := 1, 2
	fmt.Println("first: ", first, "second: ", second)
	result := Add(first, second)
	fmt.Println(first, "+", second, "=", result)
	stringResult := Add("Hello, ", "World!")
	fmt.Println("Hello, + World! = ", stringResult)
}

func Add[T int | float64 | string](a, b T) T {
	return a + b
}
