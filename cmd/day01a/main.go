package main

import (
	"fmt"
	"math"

	"github.com/scottyw/adventofcode2019/pkg/aoc"
)

func main() {
	xs := aoc.FileLinesToIntSlice("input/01.txt")
	total := 0
	for _, x := range xs {
		total += calculateFuel(x)
	}
	fmt.Println(total)
}

func calculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3.0)) - 2
}
