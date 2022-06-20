package search_test

import (
	"testing"
	"yuuzin217/algorithmBook/search"
)

func TestBinarySearch(t *testing.T) {
	for _, testCase := range []struct {
		req *request
		res *response
	}{
		{
			req: &request{
				target: "f",
				source: testSource,
			},
			res: &response{
				index: 5,
				value: "f",
			},
		},
		{
			req: &request{
				target: "u",
				source: testSource,
			},
			res: &response{
				index: 20,
				value: "u",
			},
		},
	} {
		index, value := search.BinarySearch(testCase.req.target, testCase.req.source)
		if err := check("TestBinarySearch", &response{index, value}, testCase.res); err != nil {
			t.Fatal(err)
		}
	}
	t.Log("TestBinarySearch OK")
}
