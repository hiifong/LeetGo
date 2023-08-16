package top_k_frequent_elements

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 5, 1, 1, 1, 1, 3, 4, 5, 4, 4, 3, 5, 6, 4, 6}
	topk := 5
	fmt.Println(topKFrequent(nums, topk))
}
