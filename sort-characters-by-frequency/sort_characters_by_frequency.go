package sort_characters_by_frequency

func frequencySort(s string) string {
	runes := []rune(s)
	l := len(runes)

	m := make(map[rune]int)
	for _, v := range runes {
		m[v]++
	}

	buckets := make([][]rune, l+1)
	for char, n := range m {
		i := n
		for ; n > 0; n-- {
			buckets[i] = append(buckets[i], char)
		}
	}

	result := make([]rune, l)
	n := 0
	for i := l; i >= 0; i-- {
		for _, v := range buckets[i] {
			result[n] = v
			n++
		}
	}
	return string(result)
}
