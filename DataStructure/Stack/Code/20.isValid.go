package Code

// 数据结构：Stack、Map
// 算法：记忆化搜索，数组栈基本操作

func isValid(s string) bool {
	// 二进制的奇偶判断
	if len(s) & 1 == 1 {
		return false
	}

	// 字符字典：byte，用 ''
	maps := map[byte]byte{
		')':'(',
		'}':'{',
		']':'[',
	}

	// 数组栈
	var stack []byte

	// 循环遍历字符串
	for i := 0; i < len(s); i++ {
		// 记忆化搜索
		if val, exists := maps[s[i]]; exists {
			// 运算，匹配栈顶元素
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			}

			// 匹配，出栈
			stack = stack[:len(stack)-1]
		} else {
			// 左括号，入栈
			stack = append(stack, s[i])
		}
	}

	// 返回结果
	return len(stack) == 0
}
