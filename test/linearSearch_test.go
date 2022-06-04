package algorithmBook

import (
	"testing"
	"yuuzin217/algorithmBook"
)

var dictionaryTest []string = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

type linearSearchReq struct {
	target     string
	dictionary []string
}

type linearSearchTest struct {
	request *linearSearchReq
	expect  *algorithmBook.SearchResult
}

func TestLinearSearch(t *testing.T) {

	execLinearSearch(
		&linearSearchTest{
			request: &linearSearchReq{
				target:     "a",
				dictionary: dictionaryTest,
			},
			expect: &algorithmBook.SearchResult{
				Index: 0,
				Value: "a",
			},
		}, t,
	)

	execLinearSearch(
		&linearSearchTest{
			request: &linearSearchReq{
				target:     "o",
				dictionary: dictionaryTest,
			},
			expect: &algorithmBook.SearchResult{
				Index: 14,
				Value: "o",
			},
		}, t,
	)

	t.Log("TestLinearSearch OK")

}

func execLinearSearch(lst *linearSearchTest, t *testing.T) {
	checkLinearSearch(
		algorithmBook.LinearSearch(lst.request.target, lst.request.dictionary),
		lst.expect, t,
	)
}

func checkLinearSearch(
	target *algorithmBook.SearchResult,
	expect *algorithmBook.SearchResult,
	t *testing.T,
) {
	if target.Index != expect.Index {
		t.Fatalf("%s index error, result: %d/ expect: %d,", "LinearSearch", target.Index, expect.Index)
	}
	if target.Value != expect.Value {
		t.Fatalf("%s value error, result: %s/ expect: %s,", "LinearSearch", target.Value, expect.Value)
	}
}
