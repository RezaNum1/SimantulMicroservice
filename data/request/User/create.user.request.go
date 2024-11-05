package user

type CreateUserRequest struct {
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required,min=1,max=200" json:"password"`
	NIP      string `json:"nip"`
	Phone    string `json:"phone"`
	Jabatan  string `json:"jabatan"`
	BankID   string `json:"bankId"`
	Type     int    `json:"type"`
}
