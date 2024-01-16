package model

import (
	"time"
)

type Product struct {
	ProductName      string
	PhotoUrl         string
	AvailableAtStart time.Time
	AvailableAtEnd   time.Time
	RestaurantId     string
	CategoryId       string
	DefaultModel
}

type ProductGetAll struct {
	ProductName      string     `json:"product_name"`
	PhotoUrl         *string    `json:"photo_url"`
	AvailableAtStart *time.Time `json:"available_at_start"`
	AvailableAtEnd   *time.Time `json:"available_at_end"`
	RestaurantId     string     `json:"restaurant_id"`
	CategoryName     string     `json:"category_name"`
}

type ProductCreate struct {
	ProductName      string     `json:"product_name"`
	PhotoUrl         *string    `json:"photo_url"`
	AvailableAtStart *time.Time `json:"available_at_start"`
	AvailableAtEnd   *time.Time `json:"available_at_end"`
	RestaurantId     string     `json:"restaurant_id"`
	CategoryName     string     `json:"category_name"`
}
