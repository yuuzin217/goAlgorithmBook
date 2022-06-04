package algorithmBook

func BubbleSort(nums []int) []int {
	return bubbleSort(nums)
}

func MinMaxAvg(nums []int) []int {
	min, max, avg := minMaxAvg(nums)
	return []int{min, max, avg}
}

func InsertSort(nums []int) []int {
	return insertSort(nums)
}

func LinearSearch(target string, dictionary []string) *SearchResult {
	return linearSearch(target, dictionary)
}
