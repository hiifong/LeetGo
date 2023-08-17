package non_overlapping_intervals

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 1
	end := intervals[0][1]

	for _, v := range intervals {
		start := v[0]
		if start >= end {
			count++
			end = v[1]
		}
	}
	return len(intervals) - count
}
