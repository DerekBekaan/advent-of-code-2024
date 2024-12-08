package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	expected := 2

	actual := partOne(input)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}