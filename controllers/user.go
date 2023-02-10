package controllers

import (
	"net/http"
	"regexp"
)

func newUserController() *usercontroller {

	return &usercontroller{
		userIDPattern: regexp.MustCompile(`^/users`),
	}

}

type usercontroller struct {
	userIDPattern *regexp.Regexp
}

func (uc usercontroller) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("this is from usercontroll"))

}
