package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Item struct {
	ItemID   int    `json:"id"`
	ItemName string `json:"itemname"`
	ItemType string `json:"itemtype"`
}

var items []Item

// func Api() {

// 	json, _ := json.Marshal(item)

// 	fmt.Println(string(json))

// }

func ProductsHandler(w http.ResponseWriter, r *http.Request) {

	productsJson, err := json.Marshal(items)
	// fmt.Println("hii")
	// x, _ := http.Get("http://localhost:8090/1")
	// fmt.Println(x.Body)
	switch r.Method {

	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)

	case http.MethodPost:
		var addItems Item
		// var newItem Item
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)

		}

		json.Unmarshal(data, &addItems)
		items = append(items, addItems)
		x, _ := json.Marshal(items)

		w.Write(x)

	}

}
