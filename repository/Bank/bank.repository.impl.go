package bank

import (
	"Expire/helper"
	"Expire/model"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BankRepositoryImpl struct {
	Db *gorm.DB
}

func NewBankRepositoryImpl(Db *gorm.DB) BankRepository {
	return &BankRepositoryImpl{Db: Db}
}

func (t *BankRepositoryImpl) Create(bank model.Bank) *helper.CustomError {
	result := t.Db.Create(&bank)

	if result.Error != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Creating New Report.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return nil
}

func (t *BankRepositoryImpl) GetBank(id string) (*model.Bank, *helper.CustomError) {
	var bank model.Bank

	bankId, err := uuid.Parse(id)
	result := t.Db.First(&bank, "id = ?", bankId)
	if result.Error != nil || err != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Bank",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return &bank, nil
}

func (t *BankRepositoryImpl) GetAllBank() ([]model.Bank, *helper.CustomError) {
	var banks []model.Bank
	result := t.Db.Find(&banks)
	if result.Error != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reports",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return banks, nil
}

func (t *BankRepositoryImpl) FindBankById(bankId string) (*model.Bank, *helper.CustomError) {
	var err error
	bank := model.Bank{}

	if err = t.Db.Model(model.Bank{}).First(&bank, "id = ?", strings.ToLower(bankId)).Error; err != nil {
		return nil, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}
	return &bank, nil
}
