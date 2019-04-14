package problem0057

// Interval contains Start and End
type Interval struct {
	Start int
	End   int
}

var bottomLimit, upperLimit int

// InsertInterval returns an Interval slice after inserting
func InsertInterval(nonOverlappingInterval []Interval, newInterval Interval) []Interval {

	bottomLimit = newInterval.Start
	upperLimit = newInterval.End

	res := make([]Interval, 0, len(nonOverlappingInterval))
	squash := Interval{newInterval.Start, newInterval.End}
	for _, intv := range nonOverlappingInterval {
		tmp := intv
		if withinRange(bottomLimit, intv) {
			tmp.End = upperLimit
			squash.Start = tmp.Start
		}

		// if tmp.Start >

		// if upperLimit >= tmp.End && bottomLimit >= tmp.Start && bottomLimit <= tmp.End {
		// 	tmp.End = upperLimit
		// }

		// if squash.Start > tmp.Start ||  {
		// 	continue
		// }

		// if squash.End
		res = append(res, tmp)

		// if tmp.End < squash.End && tmp.Start > squash {
		// 	continue
		// }

		// if bottomLimit >= tmp.Start {
		// 	tmp.Start = bottomLimit
		// }

	}

	return res
}

func withinRange(num int, interval Interval) bool {
	if num >= interval.Start && num <= interval.End {
		return true
	}

	return false
}
