package goldman

import "fmt"

// CountPairsWithDiff returns the number of unique k-diff pairs given an array of numbers and a target diff
func CountPairsWithDiff(numbers []int, targetDiff int) int {
	globalCounter := 0
	numFreq := make(map[int]int)

	for _, num := range numbers {
		numFreq[num]++
	}

	fmt.Printf("%+v\n", numFreq)
	numSeen := make(map[int]bool)

	if targetDiff == 0 {
		for _, v := range numFreq {
			if v > 1 {
				globalCounter++
			}

		}
	} else {
		for _, num := range numbers {
			complement := targetDiff + num
			if numSeen[num] {
				continue
			}
			globalCounter = globalCounter + numFreq[complement]
			numSeen[num] = true
		}
	}

	return globalCounter
}
