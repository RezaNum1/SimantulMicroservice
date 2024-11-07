package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Leader struct {
	gorm.Model
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string     `gorm:"type:varchar(100);not null"`
	Jabatan   string     `gorm:"type:varchar(100);not null"`
	NIP       string     `gorm:"type:varchar(100);not null"`
	Phone     string     `gorm:"type:varchar(100);not null"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
	Report    []Report
}
