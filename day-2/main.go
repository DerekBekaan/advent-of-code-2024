package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Printf("Part One: %d\n", partOne(string(input)))
	fmt.Printf("Part Two: %d\n", partTwo(string(input)))
}

func partOne(input string) int {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	safeCount := 0
	for _, line := range lines {
		readings := strings.Split(line, " ")
		readingsInts := make([]int, 0, len(readings))

		for _, reading := range readings {
			readingInt, _ := strconv.Atoi(reading)
			readingsInts = append(readingsInts, readingInt)
		}

		safe := func (l, r int) bool {
			return r - l <= 3 && r - l > 0
		}
		if readingsInts[0] > readingsInts[1] {
			safe = func (l, r int) bool {
				return l - r <= 3 && l - r > 0
			}
		}
		
		valid := true
		for i := 0; i < len(readings) - 1; i++ {
			if !safe(readingsInts[i], readingsInts[i + 1]) {
				valid = false
				break
			}
		}

		if valid {
			safeCount += 1
		}
	}

	return safeCount
}

func partTwo(input string) int {
	input = strings.ReplaceAll(input, "\r", "")
	lines := strings.Split(input, "\n")

	safeCount := 0
	for _, line := range lines {
		readings := strings.Split(line, " ")
		readingsInts := make([]int, 0, len(readings))

		for _, reading := range readings {
			readingInt, _ := strconv.Atoi(reading)
			readingsInts = append(readingsInts, readingInt)
		}

		if validReadings(readingsInts, len(readingsInts)) {
			safeCount += 1
		} else {
			
			for i := 0; i < len(readingsInts); i++ {
				r := make([]int, 0, len(readingsInts) - 1)
				
				for j := 0; j < len(readingsInts); j++ {
					if i != j {
						r = append(r, readingsInts[j])
					}
				}
				
				if validReadings(r, len(r)) {
					safeCount += 1
					break
				}
			}
		}
	}

	return safeCount
}

func validReadings(readings []int, size int) bool {
	ascends := 0
	descends := 0

	for i := 0; i < size - 1; i++ {
		diff := readings[i] - readings[i + 1]

		if diff > 0 && diff <= 3 {
			descends += 1
		} else if diff < 0 && diff >= -3 {
			ascends += 1
		}
	}

	return ascends == len(readings) - 1 || descends == len(readings) - 1 
}