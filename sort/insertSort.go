package sort

// 挿入ソート

func InsertSort(n []int) []int {
	for i := 1; i < len(n); i++ {
		for j := i - 1; j >= 0; j-- {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			} else {
				break
			}
		}
	}
	return n
}
