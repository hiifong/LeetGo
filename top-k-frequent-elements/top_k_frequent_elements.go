package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for key, v := range m {
		if len(buckets[v]) <= 0 {
			buckets[v] = make([]int, 0)
		}
		buckets[v] = append(buckets[v], key)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(buckets[i]) > 0 {
			result = append(result, buckets[i]...)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}
