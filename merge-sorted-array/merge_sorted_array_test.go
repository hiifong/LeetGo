package merge_sorted_array

import (
	"fmt"
	"testing"
)

func TestMerger(t *testing.T) {
	nums1 := []int{2, 0}
	m := 1
	nums2 := []int{1}
	n := 1
	merge(nums1, m, nums2, n)
	for index, value := range nums1 {
		fmt.Println(index, "-", value)
	}
}
