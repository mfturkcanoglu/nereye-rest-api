package model

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Title string
	Checked bool
	CreatedDate time.Time
	UpdatedDate time.Time
}