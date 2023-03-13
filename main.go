package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productsJSON := `[  
        {
			"productId": 1,
			"manufacturer": "Johns-Jenkins",
			"sku": "p5z343vdS",
			"upc": "939581000000",
			"pricePerUnit": "497.45",
			"quantityOnHand": 9703,
			"productName": "sticky note"
		  },
		  {
			"productId": 2,
			"manufacturer": "Hessel, Schimmel and Feeney",
			"sku": "i7v300kmx",
			"upc": "740979000000",
			"pricePerUnit": "282.29",
			"quantityOnHand": 9217,
			"productName": "leg warmers"
		  },
		  {
			"productId": 3,
			"manufacturer": "Swaniawski, Bartoletti and Bruen",
			"sku": "q0L657ys7",
			"upc": "111730000000",
			"pricePerUnit": "436.26",
			"quantityOnHand": 5905,
			"productName": "lamp shade"
		  }  
        ]`
	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}
	return highestID + 1
}

func findProductByID(productID int) (*Product, int) {
	for i, product := range productList {
		if product.ProductID == productID {
			return &product, i
		}
	}
	return nil, 0
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	// switch r.Method {
	// case http.MethodGet:
	productsJSON, err := json.Marshal(productList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(productsJSON)

	// case http.MethodPost:
	// 	// add a new product to the list
	// 	var newProduct Product
	// 	bodyBytes, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	err = json.Unmarshal(bodyBytes, &newProduct)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	if newProduct.ProductID != 0 {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		return
	// 	}
	// 	// if newProduct.Manufacturer == "" || newProduct.Sku == "" {
	// 	// 	w.WriteHeader(http.StatusBadRequest)
	// 	// 	return
	// 	// }
	// 	newProduct.ProductID = getNextID()
	// 	productList = append(productList, newProduct)
	// 	w.WriteHeader(http.StatusCreated)
	// 	return

	// }
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	product, listItemIndex := findProductByID(productID)
	if product == nil {
		http.Error(w, fmt.Sprintf("no product with id %d", productID), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		// get a single product using the productID
		productJSON, err := json.Marshal(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJSON)

	case http.MethodPut:
		// update a product in the list
		var updatedProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedProduct.ProductID != productID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		product = &updatedProduct
		productList[listItemIndex] = *product
		w.WriteHeader(http.StatusOK)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	// http.HandleFunc("/products", productsHandler)
	// http.HandleFunc("/products/", productHandler)
	// http.HandleFunc("/try", api.ProductsHandler)
	// http.HandleFunc("/cookie", api.SetNewCookies)
	// http.ListenAndServe(":5050", nil)

	var x = make([][]int, 0, 325)

	// data[16]

	x = append(x, []int{4, 2}, []int{2, 3}, []int{1, 2})
	x = append(x[0:2], []int{2, 4})
	x = append(x, []int{1, 2})
	// x[20] = append(x[20], 3)
	// x = append(x, []int{3, 4})

	 

}
