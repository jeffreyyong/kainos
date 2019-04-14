package problem0062

import "fmt"

func uniquePaths(m, n int) int {
	// path[i][j] represents the number of different
	// paths to the (i,j) grid
	path := [100][100]int{}

	for i := 0; i < m; i++ {
		// arrive at the grid in column 0, only one path
		// because the robot can only move right
		path[i][0] = 1
	}

	for j := 0; j < n; j++ {
		// arrive at the grid in row 0, only one path
		// because the robot can only move down
		path[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// the number of paths to the (i,j) grid,
			// equal to reach the sum of the number of paths
			// above the upper grid and the left grid.
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	fmt.Println(path)
	return path[m-1][n-1]
}
