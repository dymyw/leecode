package Code

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 数据结构：Array
// 算法：多指针算法

func removeDuplicates(nums []int) int {
	// 双指针：slow
	slow := 0

	// 明确循环移动的是 fast 指针
	for fast := 1; fast < len(nums); fast++ {
		// slow 指针联动条件
		if nums[slow] != nums[fast] {
			slow++

			// 值操作
			if nums[slow] != nums[fast] {
				nums[slow] = nums[fast]
			}
		}
	}

	// slow 最后指向的就是需要计算的结果
	return slow + 1
}
