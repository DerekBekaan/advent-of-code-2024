package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", partOne(string(input)))
	fmt.Printf("Part Two: %d", partTwo(string(input)))
}

func partOne(input string) int {
	inputs := strings.ReplaceAll(input, "\r", "")
	lists := strings.Split(inputs, "\n")

	l1 := make([]int, 0, len(lists))
	l2 := make([]int, 0, len(lists))

	for _, list := range lists {
		ls := strings.Split(list, "   ")

		i1, _ := strconv.Atoi(ls[0])
		l1 = append(l1, i1)
		
		i2, _ := strconv.Atoi(ls[1])
		l2 = append(l2, i2)
	}

	slices.Sort(l1)
	slices.Sort(l2)

	sum := 0
	for i := 0; i < len(l1); i++ {
		sum += abs(l1[i] - l2[i])
	}

	return sum
}

func partTwo(input string) int {
	inputs := strings.ReplaceAll(input, "\r", "")
	lists := strings.Split(inputs, "\n")

	l1 := make([]int, 0, len(lists))

	counts := make(map[int]int)

	for _, list := range lists {
		ls := strings.Split(list, "   ")

		i1, _ := strconv.Atoi(ls[0])
		l1 = append(l1, i1)
		
		i2, _ := strconv.Atoi(ls[1])
		counts[i2] += 1
	}

	sum := 0
	for _, i := range l1 {
		sum += i * counts[i]
	}

	return sum
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
