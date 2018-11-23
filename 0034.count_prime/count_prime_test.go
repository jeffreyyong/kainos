package problem0034

import "testing"

func TestCountPrime(t *testing.T) {
	answer := countPrime(10)
	expected := 4
	if answer != 4 {
		t.Fatalf("Expected to get %v, but got %v ", expected, answer)
	}
}
