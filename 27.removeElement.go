package dymyw_leecode

// tags: Array

func removeElement(nums []int, val int) int {
	// 双指针
	slow := 0
	// 循环
	for fast := 0; fast < len(nums); fast++ {
		// 计算
		if nums[fast] != val {
			nums[slow] = nums[fast]
			// 联动
			slow += 1
		}
	}

	// 返回某个指针状态
	return slow
}
