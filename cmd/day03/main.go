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

}
