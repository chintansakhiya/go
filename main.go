package main

import (
	"net/http"

	"github.com/chintansakhiya/api"
)

// type Try struct {
// 	mess string
// }

// func (f *Try) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Println(f)

// 	fmt.Fprintf(w, "hello %s \n%v", r.URL.Path[1:], f.mess)
// }

// func barHandlar(w http.ResponseWriter, r *http.Request) {
// 	r.URL.Path = "/home/chintan.sakhiya/Documents/go/main.go"
// 	fmt.Fprintf(w, "<h1>/home/chintan.sakhiya/Documents/go/main.go</h1>")
// 	w.Write([]byte("hello from bar"))
// }

func main() {

	// var u Try = Try{"abc"}
	// u.ServeHTTP("/xyz",&u)
	// http.Handle("/", &Try{mess: "hello from try\n"})

	// http.Handle("/mess", &Try{mess: "hello from message"})

	// http.Handle("/req", &Try{mess: "hello from request"})

	// http.HandleFunc("/home/chintan.sakhiya/Documents/go/main.go", barHandlar)
	// http.ListenAndServe(":5000", nil)
	http.HandleFunc("/1", api.ProductsHandler)
	http.ListenAndServe(":8090", nil)
}
