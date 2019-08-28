package selectionsort

import "math"

// FindSmallest returns index of the smallest element in array
func FindSmallest(elements []int) int {
	smallestIndex := 0
	smallest := elements[smallestIndex]
	for i, v := range elements {
		if v <= smallest {
			smallestIndex = i
			smallest = v
		}
	}
	return smallestIndex
}

// Sort elements using selection sort algorithm
func Sort(elements []int) []int {
	sorted := make([]int, len(elements))
	for i := range elements {
		smallest := FindSmallest(elements)
		sorted[i] = elements[smallest]
		elements[smallest] = math.MaxInt64 //replace on max value instead of deletion
	}
	return sorted
}
