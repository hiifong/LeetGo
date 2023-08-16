package sort_colors

func sortColors(nums []int) {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	r := m[0]
	for i := 0; i < r; i++ {
		nums[i] = 0
	}

	w := m[1]
	for i := r; i < r+w; i++ {
		nums[i] = 1
	}

	b := m[2]
	for i := r + w; i < r+w+b; i++ {
		nums[i] = 2
	}
}
