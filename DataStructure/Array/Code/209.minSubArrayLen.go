package Code

import "math"

// https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// 数据结构: Array
// 算法：多指针算法（循环联动）

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	// 最大数
	num, sum := math.MaxInt32, 0

	// 滑动窗口
	for r < len(nums) {
		// 计算
		sum += nums[r]
		// 循环联动
		for sum >= target {
			num = min(num, r - l + 1)
			sum -= nums[l]
			l++
		}
		// 移动
		r++
	}

	if num == math.MaxInt32 {
		return 0
	}

	return num
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
