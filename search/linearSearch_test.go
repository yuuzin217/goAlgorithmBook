package search_test

import (
	"algorithmBook/search"
	"testing"
)

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
