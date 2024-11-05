package bank

import (
	request "Expire/data/request/Bank"
	response "Expire/data/response/Bank"
	"Expire/helper"
)

type BankService interface {
	Create(bank request.CreateBankRequest) *helper.CustomError
	GetBank(id string) (*response.BankResponse, *helper.CustomError)
	GetAllBank() ([]response.BankResponse, *helper.CustomError)
}
