package main

import (
	"fmt"
	"math"

	"github.com/scottyw/adventofcode2019/pkg/aoc"
)

func main() {
	xs := aoc.FileLinesToIntSlice("input/01.txt")
	var total int
	for _, x := range xs {
		total += calculateTotalFuel(x)
	}
	fmt.Println(total)
}

func calculateTotalFuel(mass int) int {
	if mass < 9 {
		return 0
	}
	fuel := calculateFuel(mass)
	return fuel + calculateTotalFuel(fuel)
}

func calculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3.0)) - 2
}
