package problem0015

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		// Avoid adding duplicate results
		// i>0 is to prevent nums[i-1] overflow
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum < 0:
				// left has to grow bigger
				left++
			case sum > 0:
				// right has to grow bigger
				right--
			default:
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// To avoid repeated additions, both left and right need
				// to be moved to different elements.
				left, right = next(nums, left, right)

			}
		}
	}
	return res
}

func next(nums []int, left, right int) (int, int) {
	for left < right {
		switch {
		case nums[left] == nums[left+1]:
			left++
		case nums[right] == nums[right-1]:
			right--
		default:
			left++
			right--
			return left, right
		}
	}
	return left, right
}
