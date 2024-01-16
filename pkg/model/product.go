package model

import (
	"time"

	"github.com/google/uuid"
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
	ID               uuid.UUID  `json:"id"`
	ProductName      string     `json:"product_name"`
	PhotoUrl         *string    `json:"photo_url"`
	AvailableAtStart *time.Time `json:"available_at_start"`
	AvailableAtEnd   *time.Time `json:"available_at_end"`
	RestaurantId     string     `json:"restaurant_id"`
	CategoryId       string     `json:"category_id"`
}

type ProductCreate struct {
	ProductName      string     `json:"product_name"`
	PhotoUrl         *string    `json:"photo_url"`
	AvailableAtStart *time.Time `json:"available_at_start"`
	AvailableAtEnd   *time.Time `json:"available_at_end"`
	RestaurantId     string     `json:"restaurant_id"`
	CategoryId       string     `json:"category_id"`
}
