package goldman

// NthUglyNumber produces the nth ugly number
func NthUglyNumber(nth int) int {
	uglyNumbers := make([]int, nth)
	for i := 0; i < nth; i++ {
		uglyNumbers[i] = 0
	}

	uglyNumbers[0] = 1

	var i2, i3, i5 int

	nextMultipleOf2 := 2
	nextMultipleOf3 := 3
	nextMultipleOf5 := 5

	for i := 0; i < nth; i++ {
		// Choose the min value of all available multiples
		uglyNumbers[i] = min([]int{nextMultipleOf2, nextMultipleOf3, nextMultipleOf5})

		// Icrement the value of index accordingly
		if uglyNumbers[i] == nextMultipleOf2 {
			i2++
			nextMultipleOf2 = uglyNumbers[i2] * 2
		}
		if uglyNumbers[i] == nextMultipleOf3 {
			i3++
			nextMultipleOf3 = uglyNumbers[i3] * 3
		}
		if uglyNumbers[i] == nextMultipleOf5 {
			i5++
			nextMultipleOf5 = uglyNumbers[i5] * 5
		}
	}
	return uglyNumbers[len(uglyNumbers)-1]
}

func min(numbers []int) int {
	minimum := numbers[0]
	for _, num := range numbers {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}
