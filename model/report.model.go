package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ID                  *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Periode             string     `gorm:"type:varchar;not null"`
	JenisTemuan         string     `gorm:"type:varchar;not null"`
	JudulTemuan         string     `gorm:"type:varchar;not null"`
	JenisPemeriksaan    string     `gorm:"type:varchar;not null"`
	JenisKantor         string     `gorm:"type:varchar;not null"`
	PoinTemuan          string     `gorm:"type:varchar;not null"`
	RincianTemuan       string     `gorm:"type:varchar;not null"`
	RencanaTindakLanjut string     `gorm:"type:varchar;"`
	DokumenTemuan       string     `gorm:"type:varchar;"`
	TanggalPemeriksaan  time.Time  `gorm:"column:tanggalPemeriksaan;not null"`
	TargetPenyelesaian  time.Time  `gorm:"column:targetPenyelesaian;not null"`

	PoinTindakLanjut     string     `gorm:"type:varchar;"`                         //
	KomitmenTindakLanjut string     `gorm:"type:varchar;"`                         //
	DokumenTindakLanjut  string     `gorm:"type:varchar;"`                         //
	WaktuPenyelesaian    time.Time  `gorm:"column:waktuPenyelesaian;default:null"` //
	Status               int        `gorm:"type:int;default:0"`
	PrevStatus           int        `gorm:"type:int;default:0"`
	CreatedAt            *time.Time `gorm:"not null;default:now()"`
	UpdatedAt            *time.Time `gorm:"not null;default:now()"`
	BankID               uuid.UUID
	SupervisorID         uuid.UUID
	LeaderID             uuid.UUID
	Bank                 Bank
	Supervisor           Supervisor
	Leader               Leader
}
