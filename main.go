package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("user")
	dbname := os.Getenv("dbname")
	password := os.Getenv("password")

	dbURL := fmt.Sprintf("host=127.0.0.1 port=5432 user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	db, err = gorm.Open("postgres", dbURL)
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
