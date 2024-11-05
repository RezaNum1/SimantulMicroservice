package report

import (
	"Expire/model"

	"github.com/google/uuid"
)

type ExternalResponse struct {
	Id    *uuid.UUID `json:"id"`
	Name  string     `json:"name"`
	Phone string     `json:"phone"`
	Bank  model.Bank `json:"bank"`
}
