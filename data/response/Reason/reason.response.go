package reason

import (
	"github.com/google/uuid"
)

type ReasonResponse struct {
	Id          *uuid.UUID `json:"id"`
	Description string     `json:"description"`
}
