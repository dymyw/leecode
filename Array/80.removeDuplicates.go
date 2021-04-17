package Array

// tags: Array

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	// 双指针
	slow := 2

	// 循环
	for fast := 2; fast < len(nums); fast++ {
		// 计算
		if nums[fast] != nums[slow - 2] {
			// 计算
			nums[slow] = nums[fast]
			// 联动
			slow++
		}
	}

	// 返回某个指针的状态
	return slow
}
