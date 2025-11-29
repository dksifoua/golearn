package arraysandslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	expected := 15
	actual := Sum(numbers)

	if actual != expected {
		t.Errorf("Expected: %d - Actual: %d - Numbers: %v", expected, actual, numbers)
	}
}

func TestSumAll(t *testing.T) {
	a, b, c := []int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{11, 12, 13, 14, 15}
	var d []int

	expected := []int{15, 40, 65, 0}
	actual := SumAll(a, b, c, d)

	if !slices.Equal(expected, actual) {
		t.Errorf("Expected: %d - Actual: %d", expected, actual)
	}
}

func TestSumAllTails(t *testing.T) {
	a, b, c := []int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{11, 12, 13, 14, 15}
	var d []int

	expected := []int{14, 34, 54, 0}
	actual := SumAllTails(a, b, c, d)

	if !slices.Equal(expected, actual) {
		t.Errorf("Expected: %d - Actual: %d", expected, actual)
	}
}
