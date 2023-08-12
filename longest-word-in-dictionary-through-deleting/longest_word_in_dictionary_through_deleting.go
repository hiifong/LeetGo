package longest_word_in_dictionary_through_deleting

import (
	"sort"
)

func findLongestWord(s string, dictionary []string) string {
	sort.Strings(dictionary)
	longWord := ""
	for _, v := range dictionary {
		l1 := len(longWord)
		l2 := len(v)
		if l1 > l2 || l1 == l2 {
			continue
		}
		if isSubstr(s, v) {
			longWord = v
		}

	}
	return longWord
}

func isSubstr(s, target string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(target) {
		if s[i] == target[j] {
			j++
		}
		i++
	}
	return j == len(target)
}
