package main

import (
	//"errors"
	"fmt"
	"strconv"
)

func main() {

	a, perror := devide(6, 2)
	if perror == nil {
		fmt.Println("division is", a)
	} else {
		fmt.Println(a, perror)
	}

}

type CustomError struct {
	mess  string
	value int
}

func (c CustomError) Error() string {
	fmt.Println("from error")
	return c.mess + " " + strconv.Itoa(c.value)

}

func devide(x, y float64) (float64, error) {

	if y == 0 {
		return x, CustomError{"this is error", -1}

	}
	return float64(x) / float64(y), nil
}
