package model

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant struct {
	ID         uuid.UUID
	PhotoUrl   string
	SignName   string
	AddressId  uuid.UUID
	CustomerId uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	Deleted    bool
}

type RestaurantGet struct {
	PhotoUrl    string  `json:"photo_url"`
	SignName    string  `json:"sign_name"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	County      string  `json:"county"`
	District    string  `json:"district"`
	FullAddress string  `json:"full_address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type RestaurantCreate struct {
	CustomerId  string  `json:"customer_id"`
	PhotoUrl    string  `json:"photo_url"`
	SignName    string  `json:"sign_name"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	County      string  `json:"county"`
	District    string  `json:"district"`
	FullAddress string  `json:"full_address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
