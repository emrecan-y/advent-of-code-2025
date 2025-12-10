package main

import (
	"aoc2025/internal/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

var (
	dayInput []string
)

func main() {
	dayInput = utils.ReadDayInput(10)

	partOne()
	partTwo()
}

func partOne() {
	count := 0
	machines := []machine{}

	for _, line := range dayInput {
		split1 := strings.Index(line, "]")
		split2 := strings.Index(line, "{") - 1

		lightsWanted := []bool{}
		for _, lightWanted := range line[1:split1] {
			lightsWanted = append(lightsWanted, lightWanted == '#')
		}
		lightsActual := make([]bool, len(lightsWanted))

		buttonWirings := [][]int{}
		for buttonWiring := range strings.FieldsSeq(line[split1+2 : split2]) {
			stringButtonWirings := strings.Split(buttonWiring[1:len(buttonWiring)-1], ",")
			buttonWiring := []int{}
			for _, numString := range stringButtonWirings {
				num, _ := strconv.ParseInt(numString, 10, 32)
				buttonWiring = append(buttonWiring, int(num))
			}
			buttonWirings = append(buttonWirings, buttonWiring)
		}

		joltageRequriements := []int{}

		for joltageString := range strings.SplitSeq(line[split2+1:], ",") {
			joltage, _ := strconv.ParseInt(joltageString, 10, 32)
			joltageRequriements = append(joltageRequriements, int(joltage))
		}
		machines = append(machines,
			machine{lightsActual, lightsWanted, buttonWirings, joltageRequriements, []int{}})
	}

	for _, machine := range machines {
		c := getMinButtonPresses(machine)
		count += c
	}

	fmt.Printf("PartOne: %v\n", count)
}

func getMinButtonPresses(m machine) int {
	queue := []machine{m}
	head := 0
	memo := map[string]int{}

	for head < len(queue) {
		current := queue[head]
		head++

		key := boolToString(current.lightsActual)
		if best, ok := memo[key]; ok && len(current.buttonPresses) >= best {
			continue
		}
		memo[key] = len(current.buttonPresses)

		if slices.Equal(current.lightsActual, current.lightsWanted) {
			return len(current.buttonPresses)
		}

		for i, buttonPresses := range current.buttonWirings {
			next := current.Clone()
			for _, index := range buttonPresses {
				next.lightsActual[index] = !next.lightsActual[index]
			}
			next.buttonPresses = append(next.buttonPresses, i)
			queue = append(queue, next)
		}
	}

	return math.MaxInt
}

func boolToString(in []bool) string {
	s := make([]byte, len(in))
	for i, v := range in {
		if v {
			s[i] = '1'
		} else {
			s[i] = '0'
		}
	}
	return string(s)
}

func partTwo() {
	fmt.Printf("PartTwo: %d\n", 2)
}

type machine struct {
	lightsActual        []bool
	lightsWanted        []bool
	buttonWirings       [][]int
	joltageRequriements []int
	buttonPresses       []int
}

func (m machine) Clone() machine {
	la := make([]bool, len(m.lightsActual))
	copy(la, m.lightsActual)

	lw := make([]bool, len(m.lightsWanted))
	copy(lw, m.lightsWanted)

	bw := make([][]int, len(m.buttonWirings))
	for i := range m.buttonWirings {
		bw[i] = make([]int, len(m.buttonWirings[i]))
		copy(bw[i], m.buttonWirings[i])
	}

	jr := make([]int, len(m.joltageRequriements))
	copy(jr, m.joltageRequriements)

	bp := make([]int, len(m.buttonPresses))
	copy(bp, m.buttonPresses)

	return machine{
		lightsActual:        la,
		lightsWanted:        lw,
		buttonWirings:       bw,
		joltageRequriements: jr,
		buttonPresses:       bp,
	}
}
