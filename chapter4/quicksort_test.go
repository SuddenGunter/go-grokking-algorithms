package quicksort

import (
	"math"
	"testing"
)

func TestSort(t *testing.T) {
	src := []int{11, 12, 0, -1, math.MinInt64, math.MaxInt64, 10, 2, 2, 2}
	result := Sort(src)

	expected := []int{math.MinInt64, -1, 0, 2, 2, 2, 10, 11, 12, math.MaxInt64}
	for i := range expected {
		if result[i] != expected[i] {
			t.Logf("Result: %v", result)
			t.FailNow()
		}
	}
}

func Test_Sort(t *testing.T) {
	src := []int{4, 3, 1, 5, 2}
	result := Sort(src)

	expected := []int{1, 2, 3, 4, 5}
	for i := range expected {
		if result[i] != expected[i] {
			t.Logf("Result: %v", result)
			t.FailNow()
		}
	}
}
