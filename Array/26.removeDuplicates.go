package Array

// tags: Array

func removeDuplicates(nums []int) int {
	// 双指针
	slow := 0

	// 循环
	for fast := 1; fast < len(nums); fast++ {
		// 计算
		if nums[fast] != nums[slow] {
			// 联动
			slow++
			// 计算
			nums[slow] = nums[fast]
		}
	}

	// 返回某个指针的状态
	return slow + 1
}
