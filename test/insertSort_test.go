package algorithmBook

import (
	"testing"
	"yuuzin217/algorithmBook"
)

func TestInsertSort(t *testing.T) {

	label := "InsertSort"

	exec(
		[]int{4, 2, 5, 1, 3},
		[]int{1, 2, 3, 4, 5},
		label,
		algorithmBook.InsertSort,
		t,
	)

	exec(
		[]int{25, 43, 34, 61, 85, 90, 37, 12},
		[]int{12, 25, 34, 37, 43, 61, 85, 90},
		label,
		algorithmBook.InsertSort,
		t,
	)

	t.Log("TestInsertSort OK")

}
