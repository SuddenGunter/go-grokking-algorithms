package quicksort

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func Sort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	//potentially can be paralleled
	Sort(a[:left])
	Sort(a[left+1:])

	return a
}

type ConcurrentSorter struct {
	jobs    chan []int
	counter atomic.Value
	wg      sync.WaitGroup
}

func ConcurrentSort(a []int) []int {
	sorter := &ConcurrentSorter{
		jobs: make(chan []int, 1000),
		wg:   sync.WaitGroup{},
	}
	for i := 0; i < 4; i++ {
		go sorter.Worker()
	}

	sorter.wg.Add(1)
	sorter.jobs <- a
	sorter.wg.Wait()

	return a
}
func (sorter *ConcurrentSorter) Worker() {
	func(jobs chan []int) {
		for {
			select {
			case toSort := <-jobs:
				sortChunk(toSort, jobs, &sorter.wg)
			case <-time.After(100 * time.Second):
				log.Fatal("failed")
				return
			}
		}
	}(sorter.jobs)
}

func sortChunk(a []int, jobs chan<- []int, wg *sync.WaitGroup) {
	if len(a) < 2 {
		wg.Done()
		return
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	wg.Add(2)
	jobs <- a[:left]
	jobs <- a[left+1:]

	wg.Done()
}
