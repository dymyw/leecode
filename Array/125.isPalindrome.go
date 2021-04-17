package Array

// tags: Array、String
// 马拉车算法
// https://www.mayw.tech/work/tech/2019/04/02/Manacher's-Algorithm-%E9%A9%AC%E6%8B%89%E8%BD%A6%E7%AE%97%E6%B3%95.html

import "strings"

func isPalindrome(s string) bool {
	// 初始化处理
	var ns string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			ns += string(s[i])
		}
	}
	ns = strings.ToLower(ns)

	// 循环
	for i := 0; i < (len(ns) >> 1); i++ {
		// 指针碰撞
		if ns[i] != ns[len(ns)-1-i] {
			return false
		}
	}

	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
