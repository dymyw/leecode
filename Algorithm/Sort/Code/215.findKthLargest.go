package Code

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// 数据结构: Array
// 算法：快速排序

func findKthLargest(nums []int, k int) int {
	nums = QuickSort(nums)
	return nums[len(nums) - k]
}

func QuickSort(s []int) []int {
	// 递归终止条件
	if len(s) <= 1 {
		return s
	}

	// 处理当前问题
	mid := s[0] // 基准
	head, tail := 0, len(s)-1
	for i := 1; i <= tail; {
		if s[i] > mid {
			s[tail], s[i] = s[i], s[tail]
			tail--
		} else {
			s[head], s[i] = s[i], s[head]
			head++
			i++
		}
	}

	// 计算子问题
	QuickSort(s[:head])
	QuickSort(s[head+1:])

	// 返回结果
	return s
}
