package Code

// https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
// 数据结构: Array、String
// 算法：多指针算法

func reverseVowels(s string) string {
	// string to byte
	b := []byte(s)

	// 循环
	// for 的特殊写法
	for l, r := 0, len(b) - 1; l < r; {
		// 定位
		if !isVowel(b[l]) {
			l++
			// 继续
			continue
		}
		if !isVowel(b[r]) {
			r--
			// 继续
			continue
		}
		// 运算
		b[l], b[r] = b[r], b[l]
		// 联动
		l++
		r--
	}

	// byte to string
	return string(b)
}

func isVowel(s byte) bool {
	// use set ?
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' ||
		s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}

	return false
}