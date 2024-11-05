package bank

type CreateBankRequest struct {
	Name    string `validate:"required,min=1,max=200" json:"name"`
	Address string `validate:"required,min=1,max=200" json:"address"`
}
