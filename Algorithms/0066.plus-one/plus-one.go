package problem0066

func plusOne(digits []int) []int {
	size := len(digits)

	// 空数组
	if size == 0 {
		return []int{1}
	}

	digits[size-1]++
	// 末位进位
	for i := size-1; i > 0; i-- {
		if digits[i] == 10 {
			digits[i] -= 10
			digits[i-1] += 1
		}
	}

	// 首位判断
	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}

	return digits
}
