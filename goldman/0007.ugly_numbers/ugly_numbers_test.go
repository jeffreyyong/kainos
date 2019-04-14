package goldman

import "testing"

// TestNthUglyNumber tests the Nth ugly number
func TestNthUglyNumber(t *testing.T) {
	tests := []struct {
		input          int
		expectedOutput int
	}{
		{1, 1},
		// {7, 8},
		// {10, 12},
		// {15, 24},
		// {150, 5832},
	}

	for _, tt := range tests {
		output := NthUglyNumber(tt.input)

		if output != tt.expectedOutput {
			t.Fatalf("Expected output to be %v, but got %v", tt.expectedOutput, output)
		}
	}

}
