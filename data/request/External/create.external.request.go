package external

type CreateExternalRequest struct {
	Name   string `validate:"required,min=1,max=200" json:"name"`
	Phone  string `json:"phone"`
	BankID string `json:"bankId"`
}
