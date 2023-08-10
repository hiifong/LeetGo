package palindrome_number

import "fmt"

func isPalindrome(x int) bool {
	// 把数值转为字符串
	before := fmt.Sprintf("%v", x)

	// 把字符串转为 rune 切片
	runes := []rune(before)

	// 翻转 rune 切片
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// 使用 utf-8 把 run 切片转为字符串
	after := string(runes)
	if before == after {
		return true
	}
	return false
}
