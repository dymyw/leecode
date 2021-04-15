package dymyw_leecode

func isValid(s string) bool {
	// 奇偶判断
	if len(s) & 1 == 1 {
		return false
	}

	// 字符字典：byte、''
	maps := map[byte]byte{
		')':'(',
		'}':'{',
		']':'[',
	}

	// 数组栈
	stack := []byte{}
	// 循环遍历
	for i := 0; i < len(s); i++ {
		// 记忆化搜索
		if val, exists := maps[s[i]]; exists {
			// 运算
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			} else {
				// 出栈
				stack = stack[:len(stack)-1]
			}
		} else {
			// 入栈
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

