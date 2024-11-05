package response

import uuid "github.com/satori/go.uuid"

type UserResponse struct {
	Id    *uuid.UUID `json:"id"`
	Email string     `json:"email"`
}
