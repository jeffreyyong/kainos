package goldman

import "testing"

func TestIsPowerOfTen(t *testing.T) {
	tests := []struct {
		num            float64
		expectedOutput bool
	}{
		{9.0, false},
		{10.0, true},
		{100, true},
		{10000000000, true},
		{1, true},
		{0.1, true},
		{0.001, true},
	}

	for _, tt := range tests {
		output := IsPowerOfTen(tt.num)
		if output != tt.expectedOutput {
			t.Fatalf("Expected output to be %v, but got %v", tt.expectedOutput, output)
		}
	}

}
