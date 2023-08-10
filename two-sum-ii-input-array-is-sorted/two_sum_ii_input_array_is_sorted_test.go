package two_sum_ii_input_array_is_sorted

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}
