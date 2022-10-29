package search

// 二分探索法

func BinarySearch(target string, source []string) (int, string) {
	idx, low, high := 0, 0, len(source)
	for idx == 0 && low <= high {
		mid := (low + high) / 2
		if target == source[mid] {
			idx = mid
		} else {
			if target < source[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return idx, source[idx]
}
