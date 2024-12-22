package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", partOne(string(input)))
}

func partOne(input string) int {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	height := len(lines)
	width := len(lines[0])

	letterGrid := make([][]rune, height)
	for i, line := range lines {
		letterGrid[i] = make([]rune, width)

		for j, c := range line {
			letterGrid[i][j] = c
		}
	}

	directions := [][]int {
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	word := "XMAS"
	sum := 0
	for row := range letterGrid {
		for col := range letterGrid[row] {
			for _, d := range directions {
				for i, c := range word {
					x := col + d[0] * i
					y := row + d[1] * i

					if x < 0 || x >= width || y < 0 || y >= height || c != letterGrid[x][y]{
						break
					}
					
					if i == len(word) - 1 {
						sum += 1
					}
				}
			}
		}
	}

	return sum
}