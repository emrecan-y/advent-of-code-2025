package main

import (
	"aoc2025/internal/utils"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

var (
	dayInput []string
)

func main() {
	dayInput = utils.ReadDayInput(3)

	partOne()
	partTwo()
}

func partOne() {
	var sum atomic.Int64
	var wg sync.WaitGroup

	for _, line := range dayInput {
		wg.Go(func() {
			highestJoltage, _ := strconv.ParseInt(findHighestJoltage(line, 2), 10, 64)
			sum.Add(highestJoltage)
		})
	}
	wg.Wait()
	fmt.Printf("PartOne: %d\n", sum.Load())
}

func partTwo() {
	var sum atomic.Int64
	var wg sync.WaitGroup

	for _, line := range dayInput {
		wg.Go(func() {
			highestJoltage, _ := strconv.ParseInt(findHighestJoltage(line, 12), 10, 64)
			sum.Add(highestJoltage)
		})
	}
	wg.Wait()
	fmt.Printf("PartTwo: %d\n", sum.Load())
}

func findHighestJoltage(nums string, digits int8) string {
	if digits == 0 {
		return ""
	}

	searchableSubstr := nums[:len(nums)-int(digits)+1]

	highestJoltage := '0'
	highestJoltageIndex := 0
	for i, char := range searchableSubstr {
		if char > highestJoltage {
			highestJoltage = char
			highestJoltageIndex = i
		}
	}

	remainingSubstr := nums[highestJoltageIndex+1:]

	return string(highestJoltage) + findHighestJoltage(remainingSubstr, digits-1)
}
