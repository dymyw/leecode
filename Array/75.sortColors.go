package Array

// tags: Array

func sortColors(nums []int)  {
	// 函数变量
	var compare func(index int)
	// 递归
	compare = func(index int) {
		if index == 0 {
			return
		}
		if nums[index-1] > nums[index] {
			nums[index-1], nums[index] = nums[index], nums[index-1]
		}

		// 递归
		compare(index - 1)
	}

	for i := 0; i < len(nums); i++ {
		compare(i)
	}
}
