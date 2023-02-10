package controllers

import "net/http"

func RegisterController() {
	uc := newUserController()

	http.Handle("/user", *uc)
	http.Handle("/user/", *uc)
}
