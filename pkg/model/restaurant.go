package model

import (
	"time"

	"github.com/google/uuid"
)

// Restaurant Model in DB
type Restaurant struct {
	ID                      uuid.UUID
	PhotoUrl                string
	SignName                string
	AboutUs                 string
	ExtraInfo               string
	PhoneNumber             string
	WorkplacePhoneNumber    string
	IsAvailable             bool
	AvailableAtStart        time.Time
	AvailableAtEnd          time.Time
	WeekendAvailableAtStart time.Time
	WeekendAvailableAtEnd   time.Time
	AddressId               uuid.UUID
	CustomerId              uuid.UUID
	CreatedAt               time.Time
	UpdatedAt               time.Time
	DeletedAt               time.Time
	Deleted                 bool
}

type RestaurantGet struct {
	PhotoUrl                string     `json:"photo_url"`
	SignName                string     `json:"sign_name"`
	AboutUs                 *string    `json:"about_us"`
	ExtraInfo               *string    `json:"extra_info"`
	PhoneNumber             *string    `json:"phone_number"`
	WorkplacePhoneNumber    *string    `json:"workplace_phone_number"`
	IsAvailable             *bool      `json:"is_available"`
	AvailableAtStart        *time.Time `json:"available_at_start"`
	AvailableAtEnd          *time.Time `json:"available_at_end"`
	WeekendAvailableAtStart *time.Time `json:"weekend_available_at_start"`
	WeekendAvailableAtEnd   *time.Time `json:"weekend_available_at_end"`
	Country                 string     `json:"country"`
	City                    string     `json:"city"`
	County                  string     `json:"county"`
	District                string     `json:"district"`
	FullAddress             *string    `json:"full_address"`
	Latitude                *float64   `json:"latitude"`
	Longitude               *float64   `json:"longitude"`
}

type RestaurantCreate struct {
	CustomerId              string     `json:"customer_id"`
	PhotoUrl                string     `json:"photo_url"`
	SignName                string     `json:"sign_name"`
	AboutUs                 *string    `json:"about_us"`
	ExtraInfo               *string    `json:"extra_info"`
	PhoneNumber             *string    `json:"phone_number"`
	WorkplacePhoneNumber    string     `json:"workplace_phone_number"`
	IsAvailable             *bool      `json:"is_available"`
	AvailableAtStart        *time.Time `json:"available_at_start"`
	AvailableAtEnd          *time.Time `json:"available_at_end"`
	WeekendAvailableAtStart *time.Time `json:"weekend_available_at_start"`
	WeekendAvailableAtEnd   *time.Time `json:"weekend_available_at_end"`
	Country                 string     `json:"country"`
	City                    string     `json:"city"`
	County                  string     `json:"county"`
	District                string     `json:"district"`
	FullAddress             *string    `json:"full_address"`
	Latitude                *float64   `json:"latitude"`
	Longitude               *float64   `json:"longitude"`
}
