package Sort

// QuickSort 快速排序，正序
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

// QuickRSort 快速排序，逆序
func QuickRSort(s []int) []int {
	// 递归终止条件
	if len(s) <= 1 {
		return s
	}

	// 处理当前问题
	mid := s[0] // 基准
	head, tail := 0, len(s)-1
	for i := 1; i <= tail; {
		if s[i] < mid {
			s[tail], s[i] = s[i], s[tail]
			tail--
		} else {
			s[head], s[i] = s[i], s[head]
			head++
			i++
		}
	}

	// 计算子问题
	QuickRSort(s[:head])
	QuickRSort(s[head+1:])

	// 返回结果
	return s
}
