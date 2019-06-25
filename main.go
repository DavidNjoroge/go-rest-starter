package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ShopItem struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
}

type ShopItems []ShopItem

var shopItems ShopItems

func main() {
	shopItems = append(shopItems, ShopItem{Id: "1", Name: "unga", Price: "100"})
	router := mux.NewRouter()
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getItems(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(shopItems)
}

func getItem(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range shopItems {
		if item.Id == params["id"] {

			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(ShopItem{})
}

func createItem(w http.ResponseWriter, req *http.Request) {
	var shopItem ShopItem
	json.NewDecoder(req.Body).Decode(&shopItem)
	shopItems = append(shopItems, shopItem)
	json.NewEncoder(w).Encode(shopItem)

}
