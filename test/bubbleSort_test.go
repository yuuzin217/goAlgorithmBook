package algorithmBook

import (
	"testing"
	"yuuzin217/algorithmBook"
)

func TestBubbleSort(t *testing.T) {

	label := "BubbleSort"

	// 5個
	exec(
		[]int{5, 2, 1, 4, 3},
		[]int{1, 2, 3, 4, 5},
		label,
		algorithmBook.BubbleSort,
		t,
	)

	// 8個
	exec(
		[]int{12, 9, 15, 8, 13, 10, 14, 11},
		[]int{8, 9, 10, 11, 12, 13, 14, 15},
		label,
		algorithmBook.BubbleSort,
		t,
	)

	t.Log("TestBubbleSort OK")

}
