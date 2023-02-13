package main

import (
	"fmt"
)

const (
	i = iota
)

func main() {
	a := 0

	x := sum(10, a)

	fmt.Println(x)

}

func sum(i, a int) func(int) int {

	a = a + i
	if i == 0 {
		return func(int) int {
			i = a
		}
	}
	i = i - 1
	return sum(i, a)
}
