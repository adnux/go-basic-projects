package functions

import "testing"

func TestFunctions(t *testing.T) {
	t.Run("TestSumUp", SumUpTest)
	t.Run("TransformNumbers", TransformNumbersTest)
	t.Run("GetTransformerFunction", GetTransformerFunctionTest)
	t.Run("CreateTransformerFn", CreateTransformerFnTest)
	t.Run("Factorial", FactorialTest)
	t.Run("Double", DoubleTest)
	t.Run("Triple", TripleTest)
}

func SumUpTest(t *testing.T) {
	// Example test case
	first, second := 1, 2
	expected := 3

	result := SumUp(first, second)
	if result != expected {
		t.Errorf("SumUp(%d, %d) = %d; want %d", first, second, result, expected)
	}
}

func TransformNumbersTest(t *testing.T) {
	result := TransformNumbers(&[]int{1, 2, 3}, func(n int) int {
		return n * 2
	})
	expected := []int{2, 4, 6}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("TransformNumbers() = %v; want %v", result, expected)
			break
		}
	}
}

func FactorialTest(t *testing.T) {
	result := Factorial(5)
	expected := 120

	if result != expected {
		t.Errorf("Factorial(%d) = %d; want %d", 5, result, expected)
	}
}

func GetTransformerFunctionTest(t *testing.T) {
	result := GetTransformerFunction(&[]int{1, 2, 3})
	expected := func(n int) int {
		return n * 2
	}

	if result(2) != expected(2) {
		t.Errorf("GetTransformerFunction() = %v; want %v", result(2), expected(2))
	}
}

func CreateTransformerFnTest(t *testing.T) {
	result := CreateTransformerFn(10)
	expected := func(n int) int {
		return n * 10
	}

	if result(2) != expected(2) {
		t.Errorf("CreateTransformerFn() = %v; want %v", result(2), expected(2))
	}
}

func DoubleTest(t *testing.T) {
	result := Double(5)
	expected := 10

	if result != expected {
		t.Errorf("Double(%d) = %d; want %d", 5, result, expected)
	}
}

func TripleTest(t *testing.T) {
	result := Triple(5)
	expected := 15

	if result != expected {
		t.Errorf("Triple(%d) = %d; want %d", 5, result, expected)
	}
}
