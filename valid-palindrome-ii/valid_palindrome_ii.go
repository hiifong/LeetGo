package valid_palindrome_ii

func validPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return isPalindrome(runes, i, j-1) || isPalindrome(runes, i+1, j)
		}
	}
	return true
}

func isPalindrome(s []rune, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
