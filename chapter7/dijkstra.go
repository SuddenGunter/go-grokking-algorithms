package chapter7

import (
	"math"
)

func Dijkstras(graph map[string]map[string]int, costs map[string]int, parents map[string]string) (map[string]int, map[string]string) {
	processed := make(map[string]struct{})
	node := findLowestCostNode(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]
		for k, _ := range neighbors {
			newCost := cost + neighbors[k]
			if costs[k] > newCost {
				costs[k] = newCost
				parents[k] = node
			}
		}

		processed[node] = struct{}{}
		node = findLowestCostNode(costs, processed)
	}
	return costs, parents
}

func findLowestCostNode(costs map[string]int, processed map[string]struct{}) string {
	lowestCost := math.MaxInt32
	lowestCostNode := ""
	for k := range costs {
		cost := costs[k]
		_, alreadyProcessed := processed[k]
		if cost < lowestCost && !alreadyProcessed {
			lowestCost = cost
			lowestCostNode = k
		}
	}
	return lowestCostNode
}
