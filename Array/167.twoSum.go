package Array

// tags: Array、Map

func twoSum(numbers []int, target int) []int {
	// 引入字典
	maps := make(map[int]int)

	for k, v := range numbers {
		// 记忆化搜索
		if index, exists := maps[target - v]; exists {
			return []int{index, k + 1}
		} else {
			maps[v] = k + 1
		}
	}

	return []int{}
}
