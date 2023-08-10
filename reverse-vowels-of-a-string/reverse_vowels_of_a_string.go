package reverse_vowels_of_a_string

func reverseVowels(s string) string {
	if s == "" {
		return ""
	}
	runs := []rune(s)
	i, j := 0, len(runs)-1
	for i < j {
		if isVowels(runs[i]) == true && isVowels(runs[j]) {
			runs[i], runs[j] = runs[j], runs[i]
			i++
			j--
		} else if !isVowels(runs[i]) {
			i++
		} else {
			j--
		}
	}
	return string(runs)
}

func isVowels(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
