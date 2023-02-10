package controllers

import (
	"fmt"
	"net/http"
)

func RegisterController() {

	uc := newUserController()

	http.Handle("/user", *uc)
	http.Handle("/user/", *uc)

	fmt.Println("uc and *uc", uc, *uc)
}
