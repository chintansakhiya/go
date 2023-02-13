package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

type circle struct {
	x       float64
	pointer *point
}

func (i *circle) test(j float64) {

	i.x = j

}
func main() {

	u := circle{3.14, &point{5, 6}}
	// fmt.Println(*u.pointer, &u.pointer, u.pointer)

	fmt.Println(u.pointer, &u.pointer, *u.pointer)
	u.test(4.14)

	fmt.Println(u)

}
