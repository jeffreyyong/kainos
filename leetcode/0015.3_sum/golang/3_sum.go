package main

import (
	"fmt"
	"sort"
)

func main() {
	answer := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(answer)
}

func threeSum(nums []int) [][]int {
	results := [][]int{}
	n := len(nums)
	if n == 0 || n < 3 {
		return results
	}
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < n-2; i++ {
		// to handle duplicate
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				results = append(results, []int{nums[left], nums[right], nums[i]})
				left++
				right--
				for nums[left] == nums[left-1] {
					left++
				}
				for nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return results
}
