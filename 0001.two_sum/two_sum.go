package problem0001

func twoSum(nums []int, target int) []int {

	check := make(map[int]int)

	for i, num := range nums {

		if j, ok := check[target-num]; ok {
			return []int{j, i}
		}
		check[num] = i
	}
	return nil
}
