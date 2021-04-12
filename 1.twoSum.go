package dymyw_leecode

func twoSum(nums []int, target int) []int {
	// 引入字典
	maps := make(map[int]int)

	for k, v := range nums {
		// 记忆化搜索
		if index, ok := maps[target - v]; ok {
			return []int{index, k}
		} else {
			maps[v] = k
		}
	}

	return []int{}
}
