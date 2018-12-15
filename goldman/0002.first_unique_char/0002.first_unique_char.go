package goldman

// FirstUniqueChar finds the first unique char from a string
// 1. Count the occurence of each character in the string
// 2. Return the one with the count of 1.
func FirstUniqueChar(word string) string {
	charCountMap := make(map[string]int)

	for _, s := range word {
		charCountMap[string(s)]++
	}

	for _, char := range word {
		if charCountMap[string(char)] == 1 {
			return string(char)
		}
	}

	return ""
}
