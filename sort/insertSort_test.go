package sort_test

import (
	"algorithmBook/sort"
	"testing"
)

func TestInsertSort(t *testing.T) {
	cases := []struct {
		req []int
		res []int
	}{
		{
			[]int{4, 2, 5, 1, 3},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{25, 43, 34, 61, 85, 90, 37, 12},
			[]int{12, 25, 34, 37, 43, 61, 85, 90},
		},
	}
	for _, c := range cases {
		if err := check("TestInsertSort", sort.BubbleSort(c.req), c.res); err != nil {
			t.Fatal(err)
		}
	}
	t.Log("TestInsertSort OK")
}
