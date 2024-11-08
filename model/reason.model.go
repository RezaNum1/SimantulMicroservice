package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Reason struct {
	gorm.Model
	ID           *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Description  string     `gorm:"type:varchar;not null"`
	RejectedStep int        `gorm:"type:int;"`
	CreatedAt    *time.Time `gorm:"not null;default:now()"`
	UpdatedAt    *time.Time `gorm:"not null;default:now()"`
	ReportID     string
}
