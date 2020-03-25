package problem0053

// 贪心算法
func maxSubArray(nums []int) int {
	maxSum, currentMax := -1 << 31, -1 << 31
	for _, n := range nums {
		currentMax = max(currentMax + n, n)
		maxSum = max(maxSum, currentMax)
	}
	return maxSum
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
