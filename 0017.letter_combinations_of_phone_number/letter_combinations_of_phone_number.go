package problem0017

var m = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	ret := []string{""}
	// Let digits in the digits to have the chance to iterate.
	for i := 0; i < len(digits); i++ {
		temp := []string{}
		// Let each element in ret add new letters.
		for j := 0; j < len(ret); j++ {
			for k := 0; k < len(m[digits[i]][k]); k++ {
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
		}
		ret = temp
	}
	return ret
}
