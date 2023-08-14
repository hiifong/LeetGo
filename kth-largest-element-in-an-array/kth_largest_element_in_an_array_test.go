package kth_largest_element_in_an_array

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 4, 6, 7, 8, 3, 3, 5, 6, 1, 10, 324, 23, 3, 23, 44, 4, 64, 32, 56, 56, 32, 5, 56, 23, 45}
	QuickSort(arr, 0, len(arr)-1)
	for i, v := range arr {
		fmt.Printf("index: %d --> value: %d\n", i, v)
	}
}
