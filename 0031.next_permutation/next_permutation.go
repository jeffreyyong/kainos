package problem0031

import (
	"fmt"
)

func nextPermutation(nums []int) {
	left := len(nums) - 2
	for 0 <= left && nums[left] >= nums[left+1] {
		left--
	}

	// At this point nums[left+1:] is a descending sequence
	fmt.Printf("Before reverse: %v left + 1: %v\n", nums, left+1)
	reverse(nums, left+1)
	fmt.Printf("After reverse: %v\n", nums)

	// At this point nums[left+1:] is a ascending sequence

	if left == -1 {
		return
	}

	right := search(nums, left+1, nums[left])
	nums[left], nums[right] = nums[right], nums[left]
	fmt.Printf("After search: %v\n", nums)

}

func reverse(nums []int, left int) {
	right := len(nums) - 1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// return the index number of the minimum value of > target in nums[left:]
// nums[left:] is an incrementing sequence
func search(nums []int, left, target int) int {
	right := len(nums) - 1
	left--
	for left+1 < right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
