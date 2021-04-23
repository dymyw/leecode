package Array

// tags: Array、Set

func containsDuplicate(nums []int) bool {
	// 引入集合
	set := make(map[int]bool)

	for _, v := range nums {
		// 记忆化搜索
		if _, exists := set[v]; exists {
			return true
		} else {
			set[v] = true
		}
	}

	return false
}
