package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TODO struct {
	gorm.Model // Adds some metadata fields to the table
	ID	uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	Title string
	Description string
	DueDate time.Time
}