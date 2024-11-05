package external

type UpdateExternalRequest struct {
	ID     string `validate:"required,min=1,max=200" json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	BankID string `json:"bankId"`
}
