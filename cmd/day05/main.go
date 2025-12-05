package main

import (
	"aoc2025/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

var (
	dayInput []string
	ranges   = []numRange{}
)

func main() {
	dayInput = utils.ReadDayInput(5)

	partOne()
	partTwo()
}

func partOne() {
	var count uint64

	isInRanges := func(inputToCheck uint64) bool {
		for _, r := range ranges {
			if r.isNumInRange(inputToCheck) {
				return true
			}
		}
		return false
	}

	isRangeParsingDone := false
	for _, line := range dayInput {
		if len(line) == 0 {
			isRangeParsingDone = true
			continue
		}

		if isRangeParsingDone {
			num, _ := strconv.ParseUint(line, 10, 64)
			if isInRanges(num) {
				count++
			}
		} else {
			nums := strings.Split(line, "-")
			from, _ := strconv.ParseUint(nums[0], 10, 64)
			to, _ := strconv.ParseUint(nums[1], 10, 64)
			ranges = append(ranges, numRange{from: from, to: to})
		}
	}

	fmt.Printf("PartOne: %d\n", count)
}

func partTwo() {
	var count uint64

	for i, r1 := range ranges {
		for k, r2 := range ranges {
			if r1.isEqual(r2) {
				continue
			}
			if r, err := combineRanges(r1, r2); err == nil {
				ranges[i] = r
				ranges[k] = r
			}

		}
	}

	uniqueRanges := map[numRange]bool{}
	for _, r := range ranges {
		if _, ok := uniqueRanges[r]; !ok {
			uniqueRanges[r] = true
			count += r.getLength()
		}
	}

	fmt.Printf("PartTwo: %d\n", count)
}

type numRange struct {
	from uint64
	to   uint64
}

func (r *numRange) isEqual(inputToCheck numRange) bool {
	return r.from == inputToCheck.from && r.to == inputToCheck.to
}

func (r *numRange) isNumInRange(inputToCheck uint64) bool {
	return r.from <= inputToCheck && inputToCheck <= r.to
}

func (r *numRange) isRangeInRange(in numRange) bool {
	return r.from <= in.to && in.from <= r.to
}
func (r *numRange) getLength() uint64 {
	return r.to - r.from + 1
}

func combineRanges(r1 numRange, r2 numRange) (numRange, error) {
	if r1.isRangeInRange(r2) {
		return numRange{from: min(r1.from, r2.from), to: max(r1.to, r2.to)}, nil
	} else {
		return numRange{from: 0, to: 0}, fmt.Errorf("ranges do not overlap")
	}
}
