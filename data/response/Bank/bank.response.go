package bank

import (
	"github.com/google/uuid"
)

type BankResponse struct {
	Id      *uuid.UUID `json:"id"`
	Name    string     `json:"bank"`
	Address string     `json:"address"`
}
