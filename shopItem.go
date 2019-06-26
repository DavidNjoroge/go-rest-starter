package main

import "github.com/jinzhu/gorm"

type ShopItem struct {
	gorm.Model
	Name  string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
}

type ShopItems []ShopItem
