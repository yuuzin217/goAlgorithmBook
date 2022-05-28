package algorithmBook

import (
	"testing"
)

var Dictionary []string = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

func exec(request []int, response []int, label string, aFunc func([]int) []int, t *testing.T) {
	check(aFunc(request), response, label, t)
}

func check(target []int, expect []int, label string, t *testing.T) {
	for i := range target {
		if target[i] != expect[i] {
			t.Errorf("%s Error index %d, expect: %d, actual: %d", label, i, expect[i], target[i])
		}
	}
}
