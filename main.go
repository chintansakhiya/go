package main

import (
	"fmt"
	"net/http"

	"github.com/chintansakhiya/controllers"
)

func main() {
	controllers.RegisterController()
	fmt.Println("in main")
	http.ListenAndServe(":3030", nil)

}
