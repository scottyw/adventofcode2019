package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/scottyw/adventofcode2019/pkg/aoc"
)

func main() {
	paths := aoc.FileLinesToStringSlice("input/03.txt")
	fmt.Println(findClosestIntersection(paths[0], paths[1]))
}

func findClosestIntersection(first, second string) int {

	const OFFSET = 100000

	circuit := [2 * OFFSET][2 * OFFSET]bool{}

	// Follow the first path
	x := OFFSET
	y := OFFSET
	steps := strings.Split(first, ",")
	for _, step := range steps {
		dir := step[0]
		dist := aoc.Atoi(step[1:])
		switch dir {
		case 'R':
			for i := 0; i < dist; i++ {
				x++
				circuit[x][y] = true
			}
		case 'L':
			for i := 0; i < dist; i++ {
				if x == 0 {
					fmt.Println(dir, dist)
				}
				x--
				circuit[x][y] = true
			}
		case 'U':
			for i := 0; i < dist; i++ {
				y++
				circuit[x][y] = true
			}
		case 'D':
			for i := 0; i < dist; i++ {
				y--
				if y == 0 {
					fmt.Println(dir, dist)
				}
				circuit[x][y] = true
			}
		default:
			log.Panicf("unexpected direction: %v", dir)
		}
	}

	// Follow the second path
	min := int(math.MaxInt64)
	x = OFFSET
	y = OFFSET
	steps = strings.Split(second, ",")
	for _, step := range steps {
		dir := step[0]
		dist := aoc.Atoi(step[1:])
		switch dir {
		case 'R':
			for i := 0; i < dist; i++ {
				x++
				if circuit[x][y] == true {
					min = checkmin(min, x-OFFSET, y-OFFSET)
				}
			}
		case 'L':
			for i := 0; i < dist; i++ {
				x--
				if circuit[x][y] == true {
					min = checkmin(min, x-OFFSET, y-OFFSET)
				}
			}
		case 'U':
			for i := 0; i < dist; i++ {
				y++
				if circuit[x][y] == true {
					min = checkmin(min, x-OFFSET, y-OFFSET)
				}
			}
		case 'D':
			for i := 0; i < dist; i++ {
				y--
				if circuit[x][y] == true {
					min = checkmin(min, x-OFFSET, y-OFFSET)
				}
			}
		default:
			log.Panicf("unexpected direction: %v", dir)
		}
	}

	return min
}

func checkmin(min, a, b int) int {

	return int(math.Min(float64(min), math.Abs(float64(a))+math.Abs(float64(b))))

}
