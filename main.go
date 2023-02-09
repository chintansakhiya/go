package main

import (
	"fmt"
)

func main() {

	port := 8080
	isstarted := startsserver(port)
	fmt.Println(isstarted)

}

func startsserver(p int) bool {

	fmt.Println("server started on ", p)

	return true

}
