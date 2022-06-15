package search

import "fmt"

// 線形探索法

func LinearSearch(target string, source []string) (int, string) {
	source = append(source, target)
	i := 0
	for ; source[i] != target; i++ {
	}
	if i >= len(source) {
		fmt.Println("該当データなし")
		return 0, ""
	}
	fmt.Println("探索完了")
	return i, source[i]
}
