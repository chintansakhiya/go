package api

import (
	"encoding/json"
	"net/http"
)

type Item struct {
	ItemID   int    `json:"id"`
	ItemName string `json:"itemname"`
	ItemType string `json:"itemtype"`
}

var item = &Item{ItemID: 1, ItemName: "rich dad poor dad", ItemType: "book"}

// func Api() {

// 	json, _ := json.Marshal(item)

// 	fmt.Println(string(json))

// }

func ProductsHandler(w http.ResponseWriter, r *http.Request) {

	productsJson, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(productsJson)

}
