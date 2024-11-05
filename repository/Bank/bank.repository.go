package bank

import (
	"Expire/helper"
	"Expire/model"
)

type BankRepository interface {
	Create(report model.Bank) *helper.CustomError
	GetBank(id string) (*model.Bank, *helper.CustomError)
	GetAllBank() ([]model.Bank, *helper.CustomError)
	FindBankById(bankId string) (*model.Bank, *helper.CustomError)
}
