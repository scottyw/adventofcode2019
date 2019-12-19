package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/scottyw/adventofcode2019/pkg/aoc"
)

func main() {
	paths := aoc.FileToStrings("input/03.txt")
	fmt.Println(findClosestIntersection(paths[0], paths[1]))
}

func findClosestIntersection(first, second string) int {

	const OFFSET = 100000

	circuit := [2 * OFFSET][2 * OFFSET]int{}

	// Follow the first path
	x := OFFSET
	y := OFFSET
	total := 0
	steps := strings.Split(first, ",")
	for _, step := range steps {
		dir := step[0]
		dist := aoc.Atoi(step[1:])
		for i := 0; i < dist; i++ {
			switch dir {
			case 'R':
				x++
			case 'L':
				x--
			case 'U':
				y++
			case 'D':
				y--
			default:
				log.Panicf("unexpected direction: %v", dir)
			}
			total++
			if circuit[x][y] == 0 {
				circuit[x][y] = total
			}
		}
	}

	// Follow the second path
	min := int(math.MaxInt64)
	x = OFFSET
	y = OFFSET
	total = 0
	steps = strings.Split(second, ",")
	for _, step := range steps {
		dir := step[0]
		dist := aoc.Atoi(step[1:])
		for i := 0; i < dist; i++ {
			switch dir {
			case 'R':
				x++
			case 'L':
				x--
			case 'U':
				y++
			case 'D':
				y--
			default:
				log.Panicf("unexpected direction: %v", dir)
			}
			total++
			if circuit[x][y] > 0 {
				min = checkmin(min, total, circuit[x][y])
			}
		}
	}

	return min
}

func checkmin(min, a, b int) int {

	return int(math.Min(float64(min), math.Abs(float64(a))+math.Abs(float64(b))))

}
