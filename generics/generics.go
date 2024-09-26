package generics

import "fmt"

func StartGenerics() {
	first, second := 1, 2
	fmt.Println("first: ", first, "second: ", second)
	result := add(first, second)
	fmt.Println(first, "+", second, "=", result)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
