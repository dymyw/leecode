package Code

// https://leetcode-cn.com/problems/merge-sorted-array/
// 数据结构: Array
// 算法：多指针算法

func merge(nums1 []int, m int, nums2 []int, n int)  {
	// 定义尾指针
	p1, p2, tail := m - 1, n - 1, m + n - 1;
	// 当前操作值
	var cur int

	// 循环
	for p1 >= 0 || p2 >= 0 {
		// 计算
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] >= nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}

		nums1[tail] = cur
		// 移动
		tail--
	}
}
