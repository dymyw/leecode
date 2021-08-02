package Code

// https://leetcode-cn.com/problems/move-zeroes/
// 数据结构：Array
// 算法：多指针算法

func moveZeroes(nums []int)  {
	// 终止条件
	if len(nums) < 2 {
		return
	}

	// 双指针
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		// 计算
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			// 指针联动
			slow += 1
		}
	}

	for slow < len(nums) {
		nums[slow] = 0
		slow += 1
	}
}
