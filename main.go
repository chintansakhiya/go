package main

import (
	"net/http"

	"github.com/chintansakhiya/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)

}
