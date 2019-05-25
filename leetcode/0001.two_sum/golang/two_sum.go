package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		_, found := m[n]
		if found {
			return []int{m[n], i}
		}
		m[target-n] = i
	}
	return nil
}

func main() {
	indices := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(indices)
}
