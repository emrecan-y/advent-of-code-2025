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
	sum := int64(0)

	for _, line := range dayInput {
		fistBattery := '0'
		secondBattery := '0'
		for i, char := range line {
			if char > fistBattery && i < len(line)-1 {
				fistBattery = char
				secondBattery = '0'
			} else if char > secondBattery {
				secondBattery = char
			}
		}
		fmt.Println(string(fistBattery), string(secondBattery))
		num, _ := strconv.ParseInt(string(fistBattery)+string(secondBattery), 10, 64)
		sum += num
	}
	fmt.Printf("PartOne: %d\n", sum)
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
