package search_test

import (
	"fmt"
	"testing"
	"yuuzin217/algorithmBook/search"
)

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

func TestLinearSearch(t *testing.T) {
	for _, testCase := range []struct {
		req *request
		res *response
	}{
		{
			req: &request{
				target: "a",
				source: testSource,
			},
			res: &response{
				index: 0,
				value: "a",
			},
		},
		{
			req: &request{
				target: "o",
				source: testSource,
			},
			res: &response{
				index: 14,
				value: "o",
			},
		},
	} {
		index, value := search.LinearSearch(testCase.req.target, testCase.req.source)
		if err := check("TestLinearSearch", &response{index, value}, testCase.res); err != nil {
			t.Fatal(err)
		}
	}
	t.Log("TestLinearSearch OK")
}

func check(label string, result *response, expect *response) error {
	if result.index != expect.index {
		return fmt.Errorf("%s index error, result: %d / expect: %d,", "LinearSearch", result.index, expect.index)
	}
	if result.value != expect.value {
		return fmt.Errorf("%s value error, result: %s / expect: %s,", "LinearSearch", result.value, expect.value)
	}
	return nil
}
