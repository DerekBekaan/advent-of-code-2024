package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", partOne(string(input)))
}

func partOne(input string) int {
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	muls := re.FindAllString(input, -1)
	
	sum := 0
	for _, mul := range muls {
		x, _ := strconv.Atoi(mul[strings.Index(mul, "(") + 1: strings.Index(mul, ",")])
		y, _ := strconv.Atoi(mul[strings.Index(mul, ",") + 1: strings.Index(mul, ")")])

		sum += x * y
	}

	return sum
}