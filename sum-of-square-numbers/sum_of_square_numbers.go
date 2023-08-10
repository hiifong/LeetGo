package sum_of_square_numbers

import "math"

func judgeSquareSum(c int) bool {
	i := 0
	j := int(math.Sqrt(float64(c)))
	for i <= j {
		sum := i*i + j*j
		if sum == c {
			return true
		} else if sum < c {
			i++
		} else {
			j--
		}
	}
	return false
}
