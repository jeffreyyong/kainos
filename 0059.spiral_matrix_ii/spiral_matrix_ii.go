package problem0059

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	max := n * n
	next := nextFunc(n)

	for i := 1; i <= max; i++ {
		x, y := next()
		res[x][y] = i
	}

	return res
}

func nextFunc(n int) func() (int, int) {
	top, down := 0, n-1
	left, right := 0, n-1
	x, y := 0, -1
	// The direction to move
	dx, dy := 0, 1
	return func() (int, int) {
		// move one direction
		x += dx
		y += dy
		switch {
		// if keeps moving right exceeds right boundary
		// move the upper boundary down by 1
		// change the moving direction to going down
		case y+dy > right:
			top++
			dx, dy = 1, 0
		// if keeps moving down exceeds the bottom boundary
		// move the right boundary left by 1
		// change the moving direction to be going left
		case x+dx > down:
			right--
			dx, dy = 0, -1
		// if keeps moving left exceeds the left boundary
		// move the bottom boundary up by 1
		// change the moving direction to be going up
		case y+dy < left:
			down--
			dx, dy = -1, 0
		// if keeps moving up exceeds the up boundary
		// move the left boundary right by 1
		// change the moving direction to be going right
		case x+dx < top:
			left++
			dx, dy = 0, 1
		}
		return x, y
	}
}
