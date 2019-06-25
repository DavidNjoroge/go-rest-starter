package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type ShopItem struct {
	gorm.Model
	Id    string `json:"id,omitempty";gorm:"primary_key"`
	Name  string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
}

type ShopItems []ShopItem

func getItems(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var shopItems ShopItems

	db.Find(&shopItems)
	json.NewEncoder(w).Encode(shopItems)
}

func getItem(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "get item")
}

func createItem(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "get item")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	fmt.Println("Go server :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=goshop dbname=goshop password=goshop sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&ShopItem{})

	handleRequests()
}
