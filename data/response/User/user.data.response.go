package response

import uuid "github.com/satori/go.uuid"

type UserDataResponse struct {
	Id       *uuid.UUID `json:"id"`
	Email    string     `json:"email"`
	Name     string     `json:"name"`
	NIP      string     `json:"nip"`
	Jabatan  string     `json:"jabatan"`
	Type     string     `json:"type"`
	Phone    string     `json:"phone"`
	BankID   string     `json:"bankId"`
	BankName string     `json:"bankName"`
}
