package goldman

import "testing"

// TestCountPairsWithSum tests function for count pairs with sum
func TestCountPairsWithSum(t *testing.T) {
	tests := []struct {
		description    string
		targetSum      int
		numbers        []int
		expectedOutput int
	}{
		{
			"example 1",
			6,
			[]int{1, 5, 7, -1},
			2,
		},
		{
			"example 2",
			6,
			[]int{1, 5, 7, -1, 5},
			3,
		},
		{
			"example 3",
			2,
			[]int{1, 1, 1, 1},
			6,
		},
		{
			"example 4",
			11,
			[]int{10, 12, 10, 15, -1, 7, 6,
				5, 4, 2, 1, 1, 1},
			9,
		},
	}

	for _, tt := range tests {
		output := CountPairsWithSum(tt.numbers, tt.targetSum)
		if output != tt.expectedOutput {
			t.Fatalf("Expected output to be %v, but got %v", tt.expectedOutput, output)
		}
	}
}
