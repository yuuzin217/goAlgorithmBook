package search_test

import "fmt"

var testSource []string = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

type request struct {
	target string
	source []string
}

type response struct {
	index int
	value string
}

func check(label string, result *response, expect *response) error {
	if result.index != expect.index {
		return fmt.Errorf("%s index error, result: %d / expect: %d,", label, result.index, expect.index)
	}
	if result.value != expect.value {
		return fmt.Errorf("%s value error, result: %s / expect: %s,", label, result.value, expect.value)
	}
	return nil
}
