package problem0049

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(words []string) [][]string {
	if len(words) == 0 {
		return [][]string{}
	}

	anagramMap := make(map[string][]string)
	for _, word := range words {
		sortedString := sortStrings(word)
		anagramMap[sortedString] = append(anagramMap[sortedString], word)
	}
	return valuesfromMap(anagramMap)
}

func sortStrings(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func valuesfromMap(anagramMap map[string][]string) [][]string {
	anagramSlice := make([][]string, 0)
	for _, v := range anagramMap {
		anagramSlice = append(anagramSlice, v)
	}
	fmt.Printf("anagramSlice %v, lenSlice: %v\n", anagramSlice, len(anagramSlice))
	return anagramSlice
}
