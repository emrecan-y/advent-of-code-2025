package main

import (
	"aoc2025/internal/utils"
	"fmt"
)

var (
	dayInput []string
)

func main() {
	dayInput = utils.ReadDayInput(4)

	partOne()
	partTwo()
}

func partOne() {
	var count int64

	for y, line := range dayInput {
		for x, char := range line {
			if char == '@' && isAdjacentPaperCountLessThanFour(y, x) {
				count++
			}
		}
	}

	fmt.Printf("PartOne: %d\n", count)
}

func partTwo() {
	var count int64

	for {
		changeDetected := false

		for y, line := range dayInput {
			for x, char := range line {
				if char == '@' && isAdjacentPaperCountLessThanFour(y, x) {
					count++
					dayInput[y] = utils.ReplaceAtIndex(dayInput[y], 'x', x)
					changeDetected = true
				}
			}
		}

		if !changeDetected {
			break
		}
	}

	fmt.Printf("PartTwo: %d\n", count)
}

func isAdjacentPaperCountLessThanFour(yPos int, xPos int) bool {
	var count int8

	for y := yPos - 1; y <= yPos+1; y++ {
		for x := xPos - 1; x <= xPos+1; x++ {
			if x == xPos && y == yPos ||
				y < 0 || y >= len(dayInput) ||
				x < 0 || x >= len(dayInput[y]) {
				continue
			}
			if dayInput[y][x] == '@' {
				count++
			}
		}
	}

	return count < 4
}
