package Code

// https://leetcode-cn.com/problems/container-with-most-water/
// 数据结构: Array、String
// 算法：指针碰撞算法

func maxArea(height []int) int {
	l, r := 0, len(height) - 1;
	h := 0
	area := 0

	for l < r {
		// 运算
		if height[l] < height[r] {
			h = height[l]
		} else {
			h = height[r]
		}
		newArea := (r - l) * h
		if newArea > area {
			area = newArea
		}

		// 移动
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return area
}
