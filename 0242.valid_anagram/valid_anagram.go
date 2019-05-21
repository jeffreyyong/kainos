package problem0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// convert string to []rune to adapt to Unicode characters
	sr := []rune(s)
	tr := []rune(t)

	// Because []rune is used, rec has to use map
	rec := make(map[rune]int, len(sr))
	for i := range sr {
		rec[sr[i]]++
		rec[tr[i]]--
	}

	for _, n := range rec {
		if n != 0 {
			return false
		}
	}

	return true
}
