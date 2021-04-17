package dymyw_leecode

// tags: Array、String

func reverseVowels(s string) string {
	// string to byte
	b := []byte(s)

	// 循环
	// for 的特殊写法
	for l, r := 0, len(b) - 1; l < r; {
		// 定位
		for l < r && !isVowel(b[l]) {
			l++
		}
		for l < r && !isVowel(b[r]) {
			r--
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
