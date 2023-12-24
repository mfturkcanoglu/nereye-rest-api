package model

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID          uuid.UUID
	Country     string
	City        string
	County      string
	Distric     string
	FullAddress string
	Latitude    float64
	Longitude   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Deleted     bool
}
