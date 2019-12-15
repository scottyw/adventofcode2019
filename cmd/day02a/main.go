package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/scottyw/adventofcode2019/pkg/aoc"
)

func main() {
	var xs []int
	for _, s := range strings.Split(aoc.FileToString("input/02.txt"), ",") {
		xs = append(xs, aoc.Atoi(s))
	}
	xs[1] = 12
	xs[2] = 2
	fmt.Println(execute(xs, 0)[0])
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
