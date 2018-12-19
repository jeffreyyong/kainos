package goldman

// Create a map to store the frequency of each number in the array
// In the next traversal, for every element check it can be combined with any other element4
// After the completion of second traversal, we'd have twice the required value stored in counter

// CountPairsWithSum returns the number of pairs that adds up to the targetNum
func CountPairsWithSum(numbers []int, targetNum int) int {
	answerCounter := 0
	numFreq := make(map[int]int)

	// Store counts of all elements in map numFreq
	for _, num := range numbers {
		numFreq[num]++
	}

	// Iterate each elemnet and increment the count (Note that every pair is counted twice)
	for _, num := range numbers {
		complement := targetNum - num
		answerCounter = answerCounter + numFreq[complement]

		// If num and num satisfies the condition, need to ensure that count is decresed by one
		if targetNum-num == num {
			answerCounter = answerCounter - 1
		}
	}
	// Return half of the answer
	return answerCounter / 2
}
