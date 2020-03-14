package quicksort

import (
	"encoding/gob"
	"log"
	"os"
	"testing"
)

/*
func TestSort(t *testing.T) {
	src := []int{11, 12, 0, -1, math.MinInt64, math.MaxInt64, 10, 2, 2, 2}
	result := ConcurrentSort(src)

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
	result := ConcurrentSort(src)

	expected := []int{1, 2, 3, 4, 5}
	for i := range expected {
		if result[i] != expected[i] {
			t.Logf("Result: %v", result)
			t.FailNow()
		}
	}
}*/

// bench

func BenchmarkConcurrentSort(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()
	src := GetData()
	array := make([]int, 5000000)
	copy(array, src)
	jobs := make(chan []int, 10000000)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = ConcurrentSort(array, jobs)
		b.StopTimer()
		copy(array, src)
		b.StartTimer()
	}
}

func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	b.ResetTimer()
	src := GetData()
	array := make([]int, 5000000)
	copy(array, src)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_ = Sort(array)
		b.StopTimer()
		copy(array, src)
		b.StartTimer()
	}
}

func GetData() []int {
	p := make([]int, 5000000)
	file, err := os.Open("array.dat")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	return p
}
