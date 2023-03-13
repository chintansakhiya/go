package api

import (
	"fmt"
	"net/http"
	"time"
)

func SetNewCookies(w http.ResponseWriter, r *http.Request) {
	setNewCookie, err := r.Cookie("chintan")
	if err == nil {
		fmt.Println(setNewCookie)

		w.Write([]byte("cookies exisits"))

	}

	if err != nil {
		setNewCookie = &http.Cookie{
			Name:    "chintan",
			Path:    "cookies",
			Value:   "chintan sakhiya",
			Expires: time.Now().Add(time.Second * 5),
		}
		http.SetCookie(w, setNewCookie)
		x := "cookies added successfully"
		w.Write([]byte(x))

		w.Write([]byte(x))

	}

}
