package controllers

import (
	"net/http"
	"regexp"
)

type usercontroller struct {
	userIDPattern *regexp.Regexp
}

func (uc usercontroller) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("this is from usercontroll"))

}

func newUserController() *usercontroller {
	return &usercontroller{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}

}
