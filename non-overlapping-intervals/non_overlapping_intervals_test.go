package non_overlapping_intervals

import (
	"fmt"
	"testing"
)

func TestI(t *testing.T) {
	intvl := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}
	fmt.Println(eraseOverlapIntervals(intvl))
}
