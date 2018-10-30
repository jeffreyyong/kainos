package problem0022

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)
	dfs(n, n, 0, bytes, &res)
	return res
}

func dfs(left, right, idx int, bytes []byte, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
		return
	}

	// "(" Don't worry about matching issues,
	// Just add left > 0 to add directly
	if left > 0 {
		bytes[idx] = '('
		dfs(left-1, right, idx+1, bytes, res)
	}

	//When you want to add ")"
	// need left < right,
	// That is, at least one of the bytes[:idx] "(" can match this ")"
	if right > 0 && left < right {
		bytes[idx] = ')'
		dfs(left, right-1, idx+1, bytes, res)
	}
}
