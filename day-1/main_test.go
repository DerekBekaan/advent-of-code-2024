package main

import "testing"

func TestPartOne(t *testing.T) {
	lists := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 11

	actual := partOne(lists)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	lists := `3   4
4   3
2   5
1   3
3   9
3   3`
	expected := 31

	actual := partTwo(lists)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}