package two_sum

func twoSum(nums []int, target int) []int {
	// 创建一个 map 来存储数组的索引和值, map 的 key 存储值,value 存储索引
	idx := make(map[int]int)
	for i, v := range nums {
		// 判断 target - nums[i] 这个可以在 map 里吗?
		if mi, ok := idx[target-nums[i]]; ok {
			return []int{mi, i}
		}
		idx[v] = i
	}
	return nil
}
