package Code

// https://leetcode-cn.com/problems/find-pivot-index/
// 数据结构: Array

func pivotIndex(nums []int) int {
	// 初始化其它值
	sum, perSum := 0, 0
	// 计算总和
	for _, v := range nums {
		sum += v
	}

	for k, v := range nums {
		// 左和 + 当前值 + 右和 = 总和
		if perSum * 2 + v == sum {
			return k
		}
		perSum += v
	}

	return -1
}
