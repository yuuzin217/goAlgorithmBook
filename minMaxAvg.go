package algorithmBook

// 最小値、最大値、平均値を求める

func minMaxAvg(nums []int) (int, int, int) {
	total, min, max := nums[0], nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
		total += num
	}
	return min, max, total / len(nums)
}
