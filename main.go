package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t.Format("Monday, Jan 2 2006"))
}
