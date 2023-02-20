package main

import "fmt"

func main() {

	x := []int{10, 20, 30, 40, 50}

	for j := range x {
		fmt.Println(j)
	}

	fmt.Println(x)

}
