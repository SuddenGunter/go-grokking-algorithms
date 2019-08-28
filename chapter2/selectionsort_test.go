package selectionsort

import "testing"

func TestFindSmallest(t *testing.T) {
	src := []int{12, 14, 10, 1, 14, 15, 12, 17}
	result := FindSmallest(src)
	if result != 3 {
		t.Logf("Result: %v", result)
		t.FailNow()
	}
}
