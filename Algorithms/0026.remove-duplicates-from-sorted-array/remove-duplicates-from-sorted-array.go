package problem0026

// 0, 0, 1, 2, 2, 2, 3
func removeDuplicates(a []int) int {
	left, right, l := 0, 1, len(a)

	for ; right < l; right++ {
		if a[left] == a[right] {
			continue
		}
		left++
		a[left], a[right] = a[right], a[left]
	}

	return left +1
}
