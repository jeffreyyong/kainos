package problem0005

func longestPalindrome(text string) string {
	// Less than 2, must be palindrome
	if len(text) < 2 {
		return text
	}

	// begin is the first character index of the longest palindrome
	// maxLen is the length of the longest palindrome
	begin, maxLen := 0, 1

	// tempBeginning represents the **first** character index number of the palindrome
	// tail represents the **tail** character index number of the palindrome
	// i represents the index number of the first character of the palindrome
	// In every loop, start with i, use the same feature of all characters in the middle segment,
	// let tempBeginning and tail point to the beginning and end of the middle segment
	// Then expand from the `median middle section` to both sides, let tempBeginning and tail
	// point to the beginning and end of the longest palindrome centered on this `median segement`
	for i := 0; i < len(text); {
		if len(text)-i <= maxLen/2 {
			// Because i is the index number of the first character
			// Suppose the length of the longest palindrome that can be found at this time is l,
			// then i <= (len(text) - i) * 2 -1
			// if len(text) - i <= maxLen / 2
			// Then l <= maxLen - 1
			// Then l <= maxLen
			break
		}

		tempBeginning, tail := i, i
		for tail < len(text)-1 && text[tail+1] == text[tail] {
			tail++
			// After the loop is over, text[tempBeginning:tail+1] is a string of identical strings
		}

		// The first character of the `medium middle segment` of the next palindrome
		// will only be text[tail + 1]
		i = tail + 1
		for tail < len(text)-1 && tempBeginning > 0 && text[tail+1] == text[tempBeginning-1] {
			tail++
			tempBeginning--
			// After the loop ends, text[tempBeginning : tail + 1] is the longest palindrome that can be found
		}

		newLen := tail + 1 - tempBeginning

		// Update the record
		if newLen > maxLen {
			begin = tempBeginning
			maxLen = newLen
		}
	}

	return text[begin : begin+maxLen]
}
