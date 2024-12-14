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
	fmt.Printf("Part Two: %d\n", partTwo(string(input)))
}

func partOne(input string) int {
	re := regexp.MustCompile(`mul\((?P<x>[0-9]+),(?P<y>[0-9]+)\)`)
	muls := re.FindAllStringSubmatch(input, -1)
	
	sum := 0
	for _, m := range muls {
		x, _ := strconv.Atoi(m[1])
		y, _ := strconv.Atoi(m[2])

		sum += x * y
	}

	return sum
}

func partTwo(input string) int {
	re := regexp.MustCompile(`.*?mul\((?P<x>[0-9]+),(?P<y>[0-9]+)\)`)	
	muls := re.FindAllStringSubmatch(input, -1)
	
	sum := 0
	do := true
	for _, m := range muls {
		if strings.Contains(m[0], "don't") {
			do = false
			continue
		}	else if strings.Contains(m[0], "do") {
			do = true
		}	

		if do {
			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])
			sum += x * y
		}
	}

	return sum
}
