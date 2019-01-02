package problem0056

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		input          []Interval
		expectedOutput []Interval
	}{
		{
			[]Interval{
				Interval{1, 3},
				Interval{2, 6},
				Interval{8, 10},
				Interval{15, 18},
			},
			[]Interval{
				Interval{1, 6},
				Interval{8, 10},
				Interval{15, 18},
			},
		},
		{
			[]Interval{
				Interval{1, 3},
				Interval{0, 6},
			},
			[]Interval{
				Interval{0, 6},
			},
		},
		{
			[]Interval{
				Interval{1, 4},
				Interval{4, 5},
			},
			[]Interval{
				Interval{1, 5},
			},
		},
		{
			[]Interval{
				Interval{1, 4},
				Interval{2, 5},
			},
			[]Interval{
				Interval{1, 5},
			},
		},
	}

	for _, tt := range tests {
		output := MergeIntervals(tt.input)

		if !reflect.DeepEqual(output, tt.expectedOutput) {
			t.Fatalf("Expected output to be %v, but got %v", tt.expectedOutput, output)
		}
	}
}
