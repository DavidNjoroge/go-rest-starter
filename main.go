package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	// router := NewRouter()
	// http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
