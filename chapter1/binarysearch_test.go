package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	src := []int{10, 11, 12, 13, 14, 15, 16, 17}
	result, err := BinarySearch(src, 13)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if result != 3 {
		t.Logf("Result: %v", result)
		t.FailNow()
	}
}
