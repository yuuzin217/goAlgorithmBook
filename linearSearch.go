package algorithmBook

import "fmt"

// 線形探索法

type SearchResult struct {
	Index int
	Value string
}

func linearSearch(target string, dictionary []string) *SearchResult {
	dictionary = append(dictionary, target)
	i := 0
	for ; dictionary[i] != target; i++ {
	}
	if i >= len(dictionary) {
		fmt.Println("該当データなし")
		return nil
	}
	fmt.Println("探索完了")
	return &SearchResult{
		Index: i,
		Value: dictionary[i],
	}
}
