package main

import "fmt"

type shape interface {
	area() float32
}

type circle struct {
	r float32
}

type rect struct {
	h float32
	w float32
}

func (c circle) area() float32 {
	return 2 * 3.1 * c.r
}

func (r rect) area() float32 {
	return r.h * r.w
}

func main() {

	a := circle{10.2}
	b := rect{10, 5}
	c := circle{10.2}
	d := circle{10.2}
	e := circle{10.2}
	f := circle{10.2}
	g := circle{10.2}
	h := circle{10.2}

	// x := a.area()
	// y := b.area()
	// fmt.Println(x, y)
	shapes := []shape{a, b, c, d, e, f, g, h}

	for _, item := range shapes {

		fmt.Println(item.area())
	}

}
