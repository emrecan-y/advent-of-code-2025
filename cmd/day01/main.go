package main

import (
	"aoc2025/internal/utils"
	"fmt"
	"strconv"
)

var (
	dayInput []string
)

func main() {
	dayInput = utils.ReadDayInput(1)

	partOne()
	partTwo()
}

func partOne() {
	sd := safeDial{val: 50}
	count := int64(0)

	for _, line := range dayInput {
		direction := line[0:1]
		num, _ := strconv.ParseInt(line[1:], 10, 64)

		if direction == "L" {
			sd.turnLeft(num)
		} else {
			sd.turnRight(num)
		}

		if sd.val == 0 {
			count++
		}
	}

	fmt.Printf("PartOne: %d\n", count)
}

func partTwo() {
	sd := safeDial{val: 50}
	count := int64(0)

	for _, line := range dayInput {
		direction := line[0:1]
		num, _ := strconv.ParseInt(line[1:], 10, 64)

		for range num {
			if direction == "L" {
				sd.turnLeft(1)
			} else {
				sd.turnRight(1)
			}

			if sd.val == 0 {
				count++
			}
		}
	}

	fmt.Printf("PartTwo: %d\n", count)
}

type safeDial struct {
	val int64
}

func (s *safeDial) turnLeft(amount int64) {
	s.val = (s.val - amount) % 100
	if s.val < 0 {
		s.val = s.val + 100
	}
}

func (s *safeDial) turnRight(amount int64) {
	s.val = (s.val + amount) % 100
}
