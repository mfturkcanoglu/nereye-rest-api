package model

import (
	"time"

	"github.com/google/uuid"
)

type DefaultModel struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Deleted   bool
}

type DefaultGetModel struct {
}
