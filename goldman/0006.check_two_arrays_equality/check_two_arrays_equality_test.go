package goldman

import "testing"

func TestSliceEquality(t *testing.T) {
	tests := []struct {
		sliceOne       []int
		sliceTwo       []int
		expectedOutput bool
	}{
		{
			[]int{1, 2, 5, 4, 0},
			[]int{2, 4, 5, 0, 1},
			true,
		},
		{
			[]int{1, 2, 5, 4, 0, 2, 1},
			[]int{2, 4, 5, 0, 1, 1, 2},
			true,
		},
		{
			[]int{1, 7, 1},
			[]int{7, 7, 1},
			false,
		},
	}

	for _, tt := range tests {
		output := SliceEquality(tt.sliceOne, tt.sliceTwo)
		if output != tt.expectedOutput {
			t.Fatalf("Expected output to be %v, but got %v", tt.expectedOutput, output)
		}
	}
}
