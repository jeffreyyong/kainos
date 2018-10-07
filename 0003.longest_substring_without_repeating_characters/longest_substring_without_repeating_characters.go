package problem0003

func lengthOfLongestSubstring(text string) int {
	location := [256]int{}
	// location[s[i]] == j means:
	// the ith character, appears in s, so, s[j+1:i] does not contain s[i]
	// location[s[i]] == -1 means s[i] appears the first time in s[i]
	for i := range location {
		// To state that haven't seen the locations
		location[i] = -1
	}

	maxLen, left := 0, 0

	for i := 0; i < len(text); i++ {
		// means s[i] already repeats at s[left:i+1]
		if location[text[i]] >= left {
			left = location[text[i]] + i
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
		location[text[i]] = i
	}
	return maxLen
}
