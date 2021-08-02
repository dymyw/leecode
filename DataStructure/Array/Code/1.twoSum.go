package Code

// https://leetcode-cn.com/problems/two-sum/
// 数据结构：Array、Map
// 算法：记忆化搜索

func twoSum(nums []int, target int) []int {
	// 引入 map
	maps := make(map[int]int)

	// 循环条件
	for k, v := range nums {
		// 记忆化搜索 O(n)
		if index, exists := maps[target - v]; exists {
			return []int{index, k}
		}

		// 记忆
		maps[v] = k
	}

	// 返回结果
	return []int{}
}
