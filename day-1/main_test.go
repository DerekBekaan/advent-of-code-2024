package main

import "testing"

func TestSumListSixDigits(t *testing.T) {
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
