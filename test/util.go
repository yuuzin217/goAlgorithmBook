package algorithmBook

import (
	"testing"
)

func exec(request []int, response []int, label string, aFunc func([]int) []int, t *testing.T) {
	check(aFunc(request), response, label, t)
}

func check(target []int, expect []int, label string, t *testing.T) {
	for i := range target {
		if target[i] != expect[i] {
			t.Fatalf("%s error, index %d, result: %d / expect: %d", label, i, target[i], expect[i])
		}
	}
}
