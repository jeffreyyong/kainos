package problem0011

func maxArea(height []int) int {
	// Look from the both ends, ensure that the width is the maximum
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		firstHeight, secondHeight := height[i], height[j]
		minTotalHeight := min(firstHeight, secondHeight)

		area := minTotalHeight * (j - i)
		if max < area {
			max = area
		}

		// Change direction toward the possibility that the area has become larger.
		if firstHeight < secondHeight {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(firstHeight, secondHeight int) int {
	if firstHeight <= secondHeight {
		return firstHeight
	}
	return secondHeight
}
