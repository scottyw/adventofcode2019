package aoc

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

// Check panics if the error is not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Insert an int into a slice of ints
func Insert(s []int, i, value int) []int {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = value
	return s
}

// Delete an int from a slice of ints
func Delete(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

// Atoi converts a string to an int and panics if it goes wrong
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

// FileToStrings reads a file into an array of strings
func FileToStrings(filePath string) []string {
	f, err := os.Open(filePath)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	Check(scanner.Err())
	return lines
}

// FileToInts reads a file into an array of strings
func FileToInts(filePath string) []int {
	f, err := os.Open(filePath)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var ints []int
	for scanner.Scan() {
		ints = append(ints, Atoi(scanner.Text()))
	}
	Check(scanner.Err())
	return ints
}

// FileToBytes reads a file into an array of bytes
func FileToBytes(filePath string) []byte {
	bs, err := ioutil.ReadFile(filePath)
	Check(err)
	return bs
}

// FileToString reads a file into a string
func FileToString(filePath string) string {
	return string(FileToBytes(filePath))
}
