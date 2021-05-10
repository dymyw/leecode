package Search

// 条件：有序、存在上下界、能通过索引访问
// 场景：有序数组

func BinarySearch(s []int, f int) (index int, exists bool) {
	if len(s) == 1 {
		if s[0] == f {
			return 0, true
		} else {
			return 0, false
		}
	}

	// 双指针
	left, right := 0, len(s) - 1
	// 循环
	for left <= right {
		// 指针运算
		mid := (left + right) >> 1
		// 值运算
		if s[mid] < f {
			// 联动
			right = mid - 1
		} else if s[mid] > f {
			// 联动
			left = mid + 1
		} else {
			// 结果
			return mid, true
		}
	}

	// 结果
	return 0, false
}
