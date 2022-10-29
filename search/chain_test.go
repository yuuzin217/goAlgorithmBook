package search_test

import (
	"algorithmBook/search"
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {

	i, v := search.SearchChain(
		248,
		map[int]*search.Chain{
			4: {
				Data:    123,
				Pointer: 1,
			},
			1: {
				Data: 248,
			},
		},
	)
	fmt.Println(i, v)

	s := search.SetChain(123, nil)
	s = search.SetChain(248, s)
	fmt.Println(s)

}
