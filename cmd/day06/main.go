package main

import (
	"aoc2025/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

var (
	dayInput []string
)

func main() {
	dayInput = utils.ReadDayInput(6)

	partOne()
	partTwo()
}

func partOne() {
	var sum int64
	var dayInput2d [][]string

	for _, line := range dayInput {
		dayInput2d = append(dayInput2d, strings.Fields(line))
	}

	for x := 0; x < len(dayInput2d[0]); x++ {
		nums := []string{}
		operation := dayInput2d[len(dayInput2d)-1][x]
		for y := 0; y < len(dayInput2d)-1; y++ {
			nums = append(nums, dayInput2d[y][x])
		}
		sum += getOperationResult(nums, operation)
	}

	fmt.Printf("PartOne: %d\n", sum)
}

func partTwo() {
	var sum int64
	transposed := transposeStringSlice(dayInput)

	nums := []string{}
	operation := ""
	for _, line := range transposed {
		if strings.TrimSpace(line) == "" {
			sum += getOperationResult(nums, operation)
			nums = []string{}
		} else {
			if strings.HasSuffix(line, "*") || strings.HasSuffix(line, "+") {
				operation = string(line[len(line)-1])
			}
			nums = append(nums, line[:len(line)-1])
		}
	}
	sum += getOperationResult(nums, operation)

	fmt.Printf("PartTwo: %d\n", sum)
}

func getOperationResult(nums []string, operation string) int64 {
	result := int64(0)
	for _, v := range nums {
		currentNum, _ := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		if result == 0 {
			result = currentNum
		} else if operation == "*" {
			result *= currentNum
		} else if operation == "+" {
			result += currentNum
		}
	}
	return result
}

func transposeStringSlice(in []string) []string {
	if len(in) == 0 {
		return []string{}
	}

	cols := len(in[0])
	rows := len(in)

	out := make([]string, cols)
	for i := range cols {
		b := make([]byte, rows)
		for j := range rows {
			b[j] = in[j][i]
		}
		out[i] = string(b)
	}

	return out
}
