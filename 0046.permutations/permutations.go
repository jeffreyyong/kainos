package problem0046

import "fmt"

func permute(nums []int) [][]int {
	n := len(nums)
	// Vector is a possible sets of answers
	vector := make([]int, n)
	// taken[i] == true means that nums[i] already appears in the vector
	taken := make([]bool, n)

	var ans [][]int

	makePermutation(0, n, nums, vector, taken, &ans)

	return ans
}

// This way is compared to my original way.
// Increased the number of comparisons
// However, the number of times the slice is generated is reduced.
// So, the speed will be faster.
func makePermutation(cur, n int, nums, vector []int, taken []bool, ans *[][]int) {
	if cur == n {
		fmt.Println("FINAL")
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("OutTaken[i] %v, curr: %v\n", i, cur)

		if !taken[i] {
			fmt.Printf("InTaken[i] %v, curr: %v\n", i, cur)

			// Prepare to use nums[i], so take[i] == true
			taken[i] = true
			vector[cur] = nums[i]

			makePermutation(cur+1, n, nums, vector, taken, ans)

			// in the next loop
			// vector[cur] = nums[i+1]
			// So, in this loop, restore taken[i] free
			// fmt.Printf("Taken[i] %v\n", i)
			taken[i] = false
		}
	}
}
