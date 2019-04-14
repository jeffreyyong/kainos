package problem0057

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	tests := []struct {
		nonOverlappingInterval []Interval
		newInterval            Interval
		expectedOutput         []Interval
	}{
		{
			[]Interval{
				Interval{1, 3},
				Interval{6, 9},
			},
			Interval{2, 5},
			[]Interval{
				Interval{1, 5},
				Interval{6, 9},
			},
		},
		{
			[]Interval{
				Interval{1, 2},
				Interval{3, 5},
				Interval{6, 7},
				Interval{8, 10},
				Interval{12, 16},
			},
			Interval{4, 8},
			[]Interval{
				Interval{1, 2},
				Interval{3, 10},
				Interval{12, 16},
			},
		},
	}

	for _, tt := range tests {
		output := InsertInterval(tt.nonOverlappingInterval, tt.newInterval)
		if !reflect.DeepEqual(output, tt.expectedOutput) {
			t.Fatalf("Expected output to be %v, but got %v", tt.expectedOutput, output)
		}
	}
}
