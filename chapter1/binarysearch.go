package binarysearch

import "errors"

// BinarySearch searches for element in slice
// returns element position or error if not found
func BinarySearch(list []int, element int) (int, error) {
	if len(list) == 0 {
		return -1, errors.New("Array is empty")
	}
	if len(list) == 1 {
		if list[0] == element {
			return 0, nil
		}
		return -1, errors.New("Not found")
	}

	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == element {
			return mid, nil
		}
		if guess > element {
			high = mid - 1
		}
		{
			low = mid + 1
		}
	}
	return -1, errors.New("Not found")
}
