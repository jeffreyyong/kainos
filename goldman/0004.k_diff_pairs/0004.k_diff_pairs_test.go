package goldman

import "testing"

func TestKDiffPairs(t *testing.T) {
	tests := []struct {
		description    string
		targetDiff     int
		numbers        []int
		expectedOutput int
	}{
		{
			"test1",
			2,
			[]int{3, 1, 4, 1, 5},
			2,
		},
		{
			"test2",
			1,
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			"test3",
			0,
			[]int{1, 3, 1, 5, 4},
			1,
		},
	}

	for _, tt := range tests {
		output := CountPairsWithDiff(tt.numbers, tt.targetDiff)
		if output != tt.expectedOutput {
			t.Fatalf("Expected output to be %v, but got %v", tt.expectedOutput, output)
		}
	}
}
