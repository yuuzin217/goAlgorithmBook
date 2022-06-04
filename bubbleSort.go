package algorithmBook

// バブルソート

func bubbleSort(nums []int) []int {
	len := len(nums) - 1
	for count := 0; count < len; count++ {
		breakFlg := true
		for i := len; count < i; i-- {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				breakFlg = false
			}
		}
		if breakFlg {
			break
		}
	}
	return nums
}
