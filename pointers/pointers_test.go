package pointers

import "testing"

func TestGetAdulthoodYears(t *testing.T) {
	// Example test case
	age := 25
	expected := 7

	result := GetAdulthoodYears(age)
	if result != expected {
		t.Errorf("GetAdulthoodYears(%d) = %d; want %d", age, result, expected)
	}
}
