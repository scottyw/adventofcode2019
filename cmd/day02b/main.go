package main

import (
	"fmt"
	"log"
	"os"

	"github.com/scottyw/adventofcode2019/pkg/aoc"
)

func main() {
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			check(x, y, 19690720)
		}
	}
}

func check(noun, verb, expected int) {
	xs := aoc.FileToIntSlice("input/02.txt")
	xs[1] = noun
	xs[2] = verb
	if execute(xs, 0)[0] == 19690720 {
		fmt.Printf("SUCCESS %d %d\n", noun, verb)
		os.Exit(0)
	}
}

func execute(code []int, pc int) []int {
	switch code[pc] {
	case 1:
		code[code[pc+3]] = code[code[pc+2]] + code[code[pc+1]]
		return execute(code, pc+4)
	case 2:
		code[code[pc+3]] = code[code[pc+2]] * code[code[pc+1]]
		return execute(code, pc+4)
	case 99:
		return code
	default:
		log.Panicf("PANIC code %v pc %d", code, pc)
	}
	return nil
}
