package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	indexMerge := m + n - 1
	for index2 >= 0 && index1 >= 0 {
		if nums2[index2] >= nums1[index1] {
			nums1[indexMerge] = nums2[index2]
			index2--
			indexMerge--
		} else {
			nums1[indexMerge] = nums1[index1]
			indexMerge--
			index1--
		}
	}
	for index2 >= 0 {
		nums1[indexMerge] = nums2[index2]
		indexMerge--
		index2--
	}
}
