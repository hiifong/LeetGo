package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	if numbers == nil || len(numbers) == 0 {
		return make([]int, 0)
	}
	i := 0
	j := len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return make([]int, 0)
}
