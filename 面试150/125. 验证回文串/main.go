package main

import (
	"strings"
)

func main() {
	println(isPalindrome("aa"))
}

// 大写转小写
// 去除所有的非数组字母字符
// ascii 48-57  表示0-9
// 97-122 表示 a-z
func isPalindrome(s string) bool {
	s = removeInvalidByte(strings.ToLower(s))
	if len(s) == 1 || len(s) == 0 {
		return true
	}
	head := 0
	tail := len(s) - 1
	for head < tail {
		if s[head] == s[tail] {
			head++
			tail--
			continue
		}

		return false
	}
	return true
}

func removeInvalidByte(s string) string {
	builder := strings.Builder{}
	for i := range s {
		if s[i] >= 48 && s[i] <= 57 {
			builder.WriteByte(s[i])
			continue
		}

		if s[i] >= 97 && s[i] <= 122 {
			builder.WriteByte(s[i])
		}
	}

	return builder.String()
}
