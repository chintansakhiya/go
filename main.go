package main

import (
	"fmt"
)

const (
	i = iota
)

func main() {

	x := 10

	y := &x
	z := &y
	a := &z

	fmt.Println(x, y, z, *y, **z, ***a)

}
