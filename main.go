package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=goshop dbname=goshop password=goshop sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&ShopItem{})

	router := NewRouter()

	fmt.Println("server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
