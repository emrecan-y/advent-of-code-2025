package main

import (
	"aoc2025/internal/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	dayInput []string
)

func main() {
	dayInput = utils.ReadDayInput(2)

	partOne()
	partTwo()
}

func partOne() {
	ranges := strings.Split(dayInput[0], ",")
	sum := int64(0)

	for _, r := range ranges {
		ids := strings.Split(r, "-")
		id1, _ := strconv.ParseInt(ids[0], 10, 64)
		id2, _ := strconv.ParseInt(ids[1], 10, 64)

		for i := id1; i <= id2; i++ {
			id := strconv.FormatInt(i, 10)
			midIndex := utf8.RuneCountInString(id) / 2
			if id[0:midIndex] == id[midIndex:] {
				sum += i
			}
		}
	}
	fmt.Printf("PartOne: %d\n", sum)
}

func partTwo() {
	ranges := strings.Split(dayInput[0], ",")
	sum := int64(0)

	for _, r := range ranges {
		ids := strings.Split(r, "-")
		id1, _ := strconv.ParseInt(ids[0], 10, 64)
		id2, _ := strconv.ParseInt(ids[1], 10, 64)

		for i := id1; i <= id2; i++ {
			id := strconv.FormatInt(i, 10)
			if !isValidId(id) {
				sum += i
			}
		}
	}
	fmt.Printf("PartTwo: %d\n", sum)
}

// An ID is invalid if it is made only of some sequence of digits repeated at least twice.
// So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times),
// and 1111111 (1 seven times) are all invalid IDs.
func isValidId(id string) bool {
	idLength := utf8.RuneCountInString(id)

	for repCount := 2; repCount <= idLength; repCount++ {
		if idLength%repCount != 0 {
			continue
		}

		subStrLength := idLength / repCount
		pattern := id[:subStrLength]
		isRepeatingPattern := true
		for i := subStrLength; i < idLength; i += subStrLength {
			if id[i:i+subStrLength] != pattern {
				isRepeatingPattern = false
				break
			}
		}
		if isRepeatingPattern {
			return false
		}
	}
	return true
}
