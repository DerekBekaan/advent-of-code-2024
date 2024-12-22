package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	expected := 18

	actual := partOne(input)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
