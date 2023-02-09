package main

import (
	"fmt"

	"github.com/chintansakhiya/models"
)

func main() {

	u := models.User{
		ID:        2,
		Firstname: "chintan",
		Lastname:  "sakhiya",
	}
	var w models.User
	w = u
	fmt.Println(u)

	j := []models.User{}

	j = append(j, u, w)

	fmt.Println(j)

}
