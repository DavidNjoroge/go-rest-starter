package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetItems(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var shopItems ShopItems

	db.Find(&shopItems)
	json.NewEncoder(w).Encode(shopItems)
}

func GetItem(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var shopItem ShopItem
	err = db.First(&shopItem, params["id"]).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shopItem)
}

func CreateItem(w http.ResponseWriter, req *http.Request) {
	var shopItem ShopItem
	json.NewDecoder(req.Body).Decode(&shopItem)
	db.Create(&shopItem)
	json.NewEncoder(w).Encode(shopItem)
}
