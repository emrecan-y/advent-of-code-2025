package main

import (
	"aoc2025/internal/utils"
	"fmt"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

var (
	dayInput    []string
	coordinates []coordinate3d

	distanceToJunction = map[float64][]coordinate3d{}
	distancesSorted    []float64
)

func main() {
	dayInput = utils.ReadDayInput(8)
	parseCoordinates()
	calcDistancesAndSort()
	partOne()
	partTwo()
}

func parseCoordinates() {
	for _, line := range dayInput {
		coord := strings.Split(line, ",")
		x, _ := strconv.ParseInt(coord[0], 10, 32)
		y, _ := strconv.ParseInt(coord[1], 10, 32)
		z, _ := strconv.ParseInt(coord[2], 10, 32)
		coordinates = append(coordinates, coordinate3d{x: int(x), y: int(y), z: int(z)})
	}
}

func calcDistancesAndSort() {
	for _, coord1 := range coordinates {
		for _, coord2 := range coordinates {
			if coord1 == coord2 {
				continue
			}
			dist := coord1.getDistance(coord2)
			distanceToJunction[dist] = []coordinate3d{coord1, coord2}
		}
	}
	distancesSorted = slices.Sorted(maps.Keys(distanceToJunction))
}

func partOne() {
	groups := []map[coordinate3d]bool{}

	for i, dist := range distancesSorted {
		if i == len(dayInput) {
			break
		}
		connection := distanceToJunction[dist]
		box1 := connection[0]
		box2 := connection[1]

		connectJunctionBoxes(&groups, box1, box2)

	}

	counts := []int64{}
	for _, v := range groups {
		counts = append(counts, int64(len(v)))
	}
	slices.Sort(counts)

	product := int64(1)
	for i := len(counts) - 1; i > 0 && i > len(counts)-4; i-- {
		product *= counts[i]
	}

	fmt.Printf("PartOne: %v\n", product)

}

func partTwo() {
	groups := []map[coordinate3d]bool{}

	for _, dist := range distancesSorted {
		connection := distanceToJunction[dist]
		box1 := connection[0]
		box2 := connection[1]

		connectJunctionBoxes(&groups, box1, box2)

		if len(groups) == 1 && len(groups[0]) == len(dayInput) {
			fmt.Printf("PartOne: %d\n", box1.x*box2.x)
			break
		}
	}

}

func connectJunctionBoxes(groups *[]map[coordinate3d]bool, box1 coordinate3d, box2 coordinate3d) {
	i1, i2 := -1, -1
	for i, group := range *groups {
		if group[box1] {
			i1 = i
		}
		if group[box2] {
			i2 = i
		}
		if i1 != -1 && i2 != -1 {
			break
		}
	}

	if i1 == -1 && i2 == -1 {
		*groups = append(*groups, map[coordinate3d]bool{box1: true, box2: true})
	} else if i1 != -1 && i2 == -1 {
		(*groups)[i1][box2] = true
	} else if i1 == -1 && i2 != -1 {
		(*groups)[i2][box1] = true
	} else if i1 != -1 && i2 != -1 && i1 != i2 {
		(*groups)[i1][box2] = true
		(*groups)[i2][box1] = true
		maps.Copy((*groups)[i1], (*groups)[i2])
		*groups = append((*groups)[:i2], (*groups)[i2+1:]...)
	}
}

type coordinate3d struct {
	x int
	y int
	z int
}

func (c *coordinate3d) getDistance(other coordinate3d) float64 {
	x := math.Abs(float64(c.x - other.x))
	y := math.Abs(float64(c.y - other.y))
	z := math.Abs(float64(c.z - other.z))

	return math.Sqrt(x*x + y*y + z*z)
}
