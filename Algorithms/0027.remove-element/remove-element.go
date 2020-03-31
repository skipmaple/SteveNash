package problem0027

func removeElement(nums []int, val int) int {
	// left为第一个为val的索引值 right为最后一个不为val的索引值
	left, right := 0, len(nums)-1
	for {
		if left < len(nums) && nums[left] != val {
			left++
		}

		if right >= 0 && nums[right] == val {
			right--
		}

		if left > right {
			break
		}

		nums[left], nums[right] = nums[right], nums[left]
	}

	return left
}
