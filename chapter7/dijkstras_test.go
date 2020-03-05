package chapter7

import (
	"math"
	"testing"
)

func TestDijkstras(t *testing.T) {
	graph := make(map[string]map[string]int)
	graph["start"] = make(map[string]int)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = make(map[string]int)
	graph["a"]["fin"] = 1

	graph["b"] = make(map[string]int)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = make(map[string]int)

	// costs
	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = math.MaxInt32

	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	costs_res, parents_res := Dijkstras(graph, costs, parents)
	if costs_res["fin"] != 6 && parents_res["fin"] != "a" {
		t.Fatal("Dijkstras failed")
	}
}
