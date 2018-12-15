package problem0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, delta := 0, math.MaxInt64

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum < target:
				left++
				if delta > target-sum {
					delta = target - sum
					res = sum
				}
			case sum > target:
				right--
				if delta > sum-target {
					delta = sum - target
					res = sum
				}
			default:
				return sum
			}
		}
	}
	return res
}
