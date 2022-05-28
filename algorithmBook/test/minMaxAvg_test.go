package algorithmBook

import (
	"testing"
	"yuuzin217/algorithmBook/algorithmBook"
)

func TestMinMaxAvg(t *testing.T) {

	label := "MinMaxAvg"

	exec(
		[]int{4, 19, 63, -32, 11, 21, 8, 92, 41, 2, -5},
		[]int{-32, 92, 20},
		label,
		algorithmBook.MinMaxAvg,
		t,
	)

	t.Log("TestMinMaxAvg OK")

}
