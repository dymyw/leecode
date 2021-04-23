package Array

// tags: Array、BinarySearch

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	mid := (left + right) >> 1

	for left <= right {
		if nums[mid] > target {
			right = mid - 1
		} else if (nums[mid] < target) {
			left = mid + 1
		} else {
			return mid
		}

		mid = (left + right) >> 1
	}

	return mid + 1
}
