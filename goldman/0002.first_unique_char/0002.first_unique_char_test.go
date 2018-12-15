package goldman

import "testing"

func TestFirstUniqueChar(t *testing.T) {
	tests := []struct {
		input       string
		expectedOut string
	}{
		{
			"leetcode",
			"l",
		},
		{
			"loveleetcode",
			"v",
		},
		{
			"abc",
			"a",
		},
	}

	for _, tt := range tests {
		output := FirstUniqueChar(tt.input)
		if output != tt.expectedOut {
			t.Fatalf("Expected output to be %v, but got %v", tt.expectedOut, output)
		}
	}
}
