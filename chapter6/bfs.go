package chapter6

import "strings"

// BFS searches for mango seller in friends list (see chapter 6 of the book)
// returns true if mango seller was found
func BFS(friends map[string][]string) bool {
	queue := NewQueue()
	for _, v := range friends["you"] {
		queue.Enqueue(v)
	}

	for !queue.IsEmpty() {
		potential := queue.Dequeue()
		if isMangoSeller(*potential) {
			return true
		}
		for _, v := range friends[*potential] {
			queue.Enqueue(v)
		}
	}

	return false
}

func isMangoSeller(name string) bool {
	return strings.HasSuffix(name, "m")
}
