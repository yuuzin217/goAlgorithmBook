package sort_test

import (
	"testing"
	"yuuzin217/algorithmBook/sort"
)

func TestBubbleSort(t *testing.T) {
	cases := []struct {
		req []int
		res []int
	}{
		{
			[]int{5, 2, 1, 4, 3},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{12, 9, 15, 8, 13, 10, 14, 11},
			[]int{8, 9, 10, 11, 12, 13, 14, 15},
		},
	}
	for _, c := range cases {
		if err := check("TestBubbleSort", sort.BubbleSort(c.req), c.res); err != nil {
			t.Fatal(err)
		}
	}
	t.Log("TestBubbleSort OK")
}
