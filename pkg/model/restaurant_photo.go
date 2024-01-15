package model

import (
	"time"

	"github.com/google/uuid"
)

type RestaurantPhoto struct {
	ID         uuid.UUID
	Name       *string
	Url        string
	ParentID   string
	PhotoOrder *int16
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	Deleted    bool
}

type RestaurantPhotoCreate struct {
	Name       *string `json:"name"`
	Url        string  `json:"url"`
	ParentID   string  `json:"parent_id"`
	PhotoOrder *int16  `json:"photo_order"`
}

type RestaurantPhotoGet struct {
	Name       *string `json:"name"`
	Url        string  `json:"url"`
	ParentID   string  `json:"parent_id"`
	PhotoOrder *int16  `json:"photo_order"`
}
