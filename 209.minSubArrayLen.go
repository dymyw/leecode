package dymyw_leecode

import "math"

// tags: Array

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	// 最大数
	num, sum := math.MaxInt32, 0

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
