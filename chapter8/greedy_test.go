package chapter8

import (
	"testing"
)

func TestGreedy(t *testing.T) {
	statesNeeded := make(map[string]struct{})
	statesNeeded["mt"] = struct{}{}
	statesNeeded["wa"] = struct{}{}
	statesNeeded["or"] = struct{}{}
	statesNeeded["id"] = struct{}{}
	statesNeeded["nv"] = struct{}{}
	statesNeeded["ut"] = struct{}{}
	statesNeeded["ca"] = struct{}{}
	statesNeeded["az"] = struct{}{}

	stations := make(map[string]map[string]struct{})
	stations["kone"] = map[string]struct{}{}
	stations["kone"]["id"] = struct{}{}
	stations["kone"]["nv"] = struct{}{}
	stations["kone"]["ut"] = struct{}{}

	stations["ktwo"] = map[string]struct{}{}
	stations["ktwo"]["wa"] = struct{}{}
	stations["ktwo"]["id"] = struct{}{}
	stations["ktwo"]["mt"] = struct{}{}

	stations["kthree"] = map[string]struct{}{}
	stations["kthree"]["or"] = struct{}{}
	stations["kthree"]["nv"] = struct{}{}
	stations["kthree"]["ca"] = struct{}{}

	stations["kfour"] = map[string]struct{}{}
	stations["kfour"]["nv"] = struct{}{}
	stations["kfour"]["ut"] = struct{}{}

	stations["kfive"] = map[string]struct{}{}
	stations["kfive"]["ca"] = struct{}{}
	stations["kfive"]["az"] = struct{}{}

	expected := map[string]struct{}{}
	expected["ktwo"] = struct{}{}
	expected["kthree"] = struct{}{}
	expected["kone"] = struct{}{}
	expected["kfive"] = struct{}{}

	result := SearchRadioCoverage(statesNeeded, stations)

	if len(expected) != len(result) {
		t.Fatalf("expected len differs from result\n")
	}
	for k, _ := range result {
		if _, exists := expected[k]; !exists {
			t.Fatalf("unexpected result %v\n", k)
		}
	}
}
