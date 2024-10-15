package lists

import "testing"

func TestLists(t *testing.T) {
	t.Run("ReplaceSliceElement", ReplaceSliceElementTest)
	t.Run("CutSlice", CutSliceTest)
}

func ReplaceSliceElementTest(t *testing.T) {
	list := []float64{1, 2, 3}
	index := 1
	value := 10.0
	expected := []float64{1, 10, 3}

	result := ReplaceSliceElement(list, index, value)
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("ReplaceSliceElement(%v, %d, %f) = %v; want %v", list, index, value, result, expected)
			break
		}
	}
}

func CutSliceTest(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5}
	start := 1
	end := 4
	expected := []float64{2, 3, 4}

	result := CutSlice(list, start, end)
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("CutSlice(%v, %d, %d) = %v; want %v", list, start, end, result, expected)
			break
		}
	}
}
