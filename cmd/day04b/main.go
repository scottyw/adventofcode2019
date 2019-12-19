package main

import "fmt"

func main() {
	total := 0
	for i := 264360; i <= 746325; i++ {
		if valid(i) {
			total++
		}
	}
	fmt.Println(total)
}

func valid(p int) bool {

	a := p / 100000 % 10
	b := p / 10000 % 10
	c := p / 1000 % 10
	d := p / 100 % 10
	e := p / 10 % 10
	f := p % 10

	return a <= b && b <= c && c <= d && d <= e && e <= f &&
		((a == b && b != c) ||
			(b == c && a != b && c != d) ||
			(c == d && b != c && d != e) ||
			(d == e && c != d && e != f) ||
			(e == f && d != e))

}
