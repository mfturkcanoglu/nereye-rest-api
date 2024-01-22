package model

import (
	"github.com/google/uuid"
)

type Product struct {
	ProductName      string
	PhotoUrl         string
	AvailableAtStart string
	AvailableAtEnd   string
	RestaurantId     string
	CategoryId       string
	DefaultModel
}

type ProductGetAll struct {
	ID               uuid.UUID `json:"id"`
	ProductName      string    `json:"product_name"`
	PhotoUrl         *string   `json:"photo_url"`
	IsAvailable      bool      `json:"is_available"`
	AvailableAtStart *string   `json:"available_at_start"`
	AvailableAtEnd   *string   `json:"available_at_end"`
	RestaurantId     string    `json:"restaurant_id"`
	CategoryId       string    `json:"category_id"`
}

type ProductCreate struct {
	ProductName      string  `json:"product_name"`
	PhotoUrl         *string `json:"photo_url"`
	AvailableAtStart *string `json:"available_at_start"`
	AvailableAtEnd   *string `json:"available_at_end"`
	RestaurantId     string  `json:"restaurant_id"`
	CategoryId       string  `json:"category_id"`
}
