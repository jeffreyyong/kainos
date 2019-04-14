package goldman

import (
	"math"
)

// IsPowerOfTen tells if the number if a power of ten
func IsPowerOfTen(num float64) bool {

	if num == 1 {
		return true
	}

	if num > 1.0 && num < 10.0 {
		return false
	}

	if math.Mod(num, 10.0) == 0 {
		return true
	}

	if math.Mod(1/num, 10) == 0 {
		return true
	}

	return false
}
