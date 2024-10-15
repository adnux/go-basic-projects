package generics

import "testing"

func TestAdd(t *testing.T) {
	t.Run("AddIntegers", AddIntegersTest)
	t.Run("AddStrings", AddStringsTest)
}

func AddIntegersTest(t *testing.T) {
	// Example test case
	first, second := 1, 2
	expected := 3

	result := Add(first, second)
	if result != expected {
		t.Errorf("Add(%d, %d) = %d; want %d", first, second, result, expected)
	}
}

func AddStringsTest(t *testing.T) {
	first, second := "Hello, ", "World!"
	expected := "Hello, World!"

	result := Add(first, second)
	if result != expected {
		t.Errorf("Add(%s, %s) = %s; want %s", first, second, result, expected)
	}
}
