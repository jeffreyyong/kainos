package goldman

// SliceEquality checks if two slices are equal, regardless of arrangement
// If there are repetitions, then counts of repeated elements must also be the same for two slices
func SliceEquality(sliceOne, sliceTwo []int) bool {
	numFreq := make(map[int]int)

	for _, num := range sliceOne {
		numFreq[num]++
	}

	for _, num := range sliceTwo {
		numFreq[num]--
	}

	for _, v := range numFreq {
		if v < 0 {
			return false
		}
	}

	return true
}
