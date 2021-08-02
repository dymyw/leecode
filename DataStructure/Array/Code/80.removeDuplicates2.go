package Code

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/
// 数据结构：Array
// 算法：多指针算法

func removeDuplicates2(nums []int) int {
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
