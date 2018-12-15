package problem0007

import (
	"math"
)

func reverse(num int) int {
	sign := 1

	if num < 0 {
		sign = -1
		num = -1 * num
	}

	result := 0
	for num > 0 {
		// Take the end of num
		temp := num % 10
		// Put at the top of result
		result = result*10 + temp
		// Num, remove the end
		num = num / 10
	}

	// Return the sign back to the result
	result = sign * result

	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}

	return result
}
