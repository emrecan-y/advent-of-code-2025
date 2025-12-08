package main

import (
	"aoc2025/internal/utils"
	"fmt"
)

var (
	dayInput []string
	memo     = map[coordinate]int64{}
)

func main() {
	dayInput = utils.ReadDayInput(7)

	partOne()
	partTwo()
}

func partOne() {
	var count int64

	for y := 0; y < len(dayInput)-1; y++ {
		for x, char := range dayInput[y] {
			if char == 'S' || char == '|' {
				if dayInput[y+1][x] == '^' {
					dayInput[y+1] = utils.ReplaceAtIndex(dayInput[y+1], '|', x+1)
					dayInput[y+1] = utils.ReplaceAtIndex(dayInput[y+1], '|', x-1)
					count++
				} else {
					dayInput[y+1] = utils.ReplaceAtIndex(dayInput[y+1], '|', x)
				}
			}
		}
	}

	fmt.Printf("PartOne: %d\n", count)
}

func partTwo() {
	startPoint := coordinate{x: len(dayInput[0]) / 2, y: 0}
	fmt.Printf("PartTwo: %d\n", dfsCount(startPoint))
}

func dfsCount(c coordinate) int64 {
	if c.y >= len(dayInput)-1 {
		return 1
	}

	char := dayInput[c.y+1][c.x]
	switch char {
	case '|':
		return dfsCount(coordinate{y: c.y + 1, x: c.x})
	case '^':
		if val, ok := memo[c]; ok {
			return val
		} else {
			leftCount := dfsCount(coordinate{y: c.y + 1, x: c.x - 1})
			rightCount := dfsCount(coordinate{y: c.y + 1, x: c.x + 1})
			sum := leftCount + rightCount
			memo[c] = sum
			return sum
		}
	default:
		return 0
	}
}

type coordinate struct {
	x int
	y int
}
