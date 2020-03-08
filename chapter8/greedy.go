package chapter8

func SearchRadioCoverage(statesNeeded map[string]struct{}, stations map[string]map[string]struct{}) map[string]struct{} {
	finalStations := make(map[string]struct{})
	for len(statesNeeded) > 0 {
		var bestStation string
		statesCovered := make(map[string]struct{})
		for station, statesCoveredByStation := range stations {
			covered := setIntersection(statesNeeded, statesCoveredByStation)
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}
		for k, _ := range statesCovered {
			delete(statesNeeded, k)
		}
		finalStations[bestStation] = struct{}{}
	}
	return finalStations
}

func setIntersection(first, second map[string]struct{}) map[string]struct{} {
	result := make(map[string]struct{})
	for k, _ := range first {
		if _, exists := second[k]; exists {
			result[k] = struct{}{}
		}
	}
	return result
}
