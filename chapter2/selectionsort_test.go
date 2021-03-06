package selectionsort

import (
	"math"
	"testing"
)

func TestFindSmallest(t *testing.T) {
	src := []int{12, 14, 10, 1, 14, 15, 12, 17}
	result := FindSmallest(src)
	if result != 3 {
		t.Logf("Result: %v", result)
		t.FailNow()
	}
}

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
