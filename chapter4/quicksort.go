package quicksort

import (
	"context"
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
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

func ConcurrentSort(a []int, c chan []int) []int {
	sorter := &ConcurrentSorter{
		jobs: c,
		wg:   sync.WaitGroup{},
	}
	shutdownWg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 4; i++ {
		shutdownWg.Add(1)
		go sorter.Worker(ctx, &shutdownWg)
	}

	sorter.wg.Add(1)
	sorter.jobs <- a
	sorter.wg.Wait()
	cancel()
	shutdownWg.Wait()
	return a
}
func (sorter *ConcurrentSorter) Worker(ctx context.Context, wg *sync.WaitGroup) {
	func(jobs chan []int) {
		for {
			select {
			case toSort := <-jobs:
				sortChunk(toSort, jobs, &sorter.wg)
			case <-ctx.Done():
				wg.Done()
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
	if len(a) < 2000 {
		sort.Ints(a)
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
