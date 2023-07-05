package problem0001

//func twoSum(nums []int, target int) []int {
//	index := make(map[int]int, len(nums))
//
//	for i,v := range nums {
//		if j, ok := index[target-v]; ok {
//			return []int {j, i}
//		}
//		index[v] = i
//	}
//
//	return nil
//}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}

	return nil
}
