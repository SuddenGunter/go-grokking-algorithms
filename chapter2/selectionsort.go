package selectionsort

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
