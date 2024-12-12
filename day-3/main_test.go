package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	// input := `mul(2,4)`

	expected := 161
	actual := partOne(input)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}