package problem0056

import "sort"

// MergeIntervals merge intervals given a slice of slices.
func MergeIntervals(input []Interval) []Interval {
	if len(input) < 2 {
		return input
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i].Start < input[j].Start
	})

	res := make([]Interval, 0, len(input))

	tmp := input[0]

	for i := 1; i < len(input); i++ {
		if tmp.End >= input[i].Start {
			tmp.End = max(input[i].End, tmp.End)
			// res = append(res, tmp)
		} else {
			res = append(res, tmp)
			tmp = input[i]
		}
	}

	res = append(res, tmp)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Interval Definition for an interval
type Interval struct {
	Start int
	End   int
}
