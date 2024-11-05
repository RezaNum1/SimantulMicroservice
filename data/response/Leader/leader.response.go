package report

import (
	"github.com/google/uuid"
)

type LeaderResponse struct {
	Id      *uuid.UUID `json:"id"`
	Name    string     `json:"name"`
	Jabatan string     `json:"jabatan"`
	NIP     string     `json:"nip"`
	Phone   string     `json:"phone"`
}
