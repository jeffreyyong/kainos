package problem0048

import (
	"fmt"
)

func rotate(m [][]int) {
	n := len(m)
	fmt.Printf("Before: %v, length: %v\n", m, n)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := m[i][j]
			// The left row is equal to the right column
			// For the case of [[1,2],[3,4]]
			// 0,0 = 1,0
			m[i][j] = m[n-j-1][i]
			// 1,0 = 1,1
			m[n-j-1][i] = m[n-i-1][n-j-1]
			// 1,1 = 0,1
			m[n-i-1][n-j-1] = m[j][n-i-1]
			// 0,1 = 0,0
			m[j][n-i-1] = temp
		}
	}
	fmt.Printf("After: %v\n", m)
}
