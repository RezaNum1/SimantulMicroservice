package bank

import (
	request "Expire/data/request/Bank"
	response "Expire/data/response/Bank"
	"Expire/helper"
	"Expire/model"
	repository "Expire/repository/Bank"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type BankServiceImpl struct {
	BankRepository repository.BankRepository
	Validate       *validator.Validate
}

func NewBankServiceImpl(bankRepository repository.BankRepository, validate *validator.Validate) BankService {
	return &BankServiceImpl{
		BankRepository: bankRepository,
		Validate:       validate,
	}
}

func (t BankServiceImpl) Create(bank request.CreateBankRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(bank)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	bankModel := model.Bank{
		Name:    bank.Name,
		Address: bank.Address,
	}
	createError := t.BankRepository.Create(bankModel)

	if createError != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return nil
}

func (t BankServiceImpl) GetBank(id string) (*response.BankResponse, *helper.CustomError) {
	result, fetchError := t.BankRepository.GetBank(id)

	if fetchError != nil {
		return nil, fetchError
	} else {
		response := response.BankResponse{
			Id:   (*uuid.UUID)(result.ID),
			Name: result.Name,
		}

		return &response, nil
	}
}

func (t BankServiceImpl) GetAllBank() ([]response.BankResponse, *helper.CustomError) {
	result, fetchError := t.BankRepository.GetAllBank()

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapBanksToBankResponse(result), nil
	}
}

func (t BankServiceImpl) mapBanksToBankResponse(banks []model.Bank) []response.BankResponse {
	responseBanks := make([]response.BankResponse, len(banks))
	for i, bank := range banks {
		responseBanks[i] = t.convertBankToBankResponse(bank)
	}
	return responseBanks
}

func (t BankServiceImpl) convertBankToBankResponse(bank model.Bank) response.BankResponse {
	// Perform necessary conversion logic here, potentially selecting specific fields
	responseBank := response.BankResponse{
		Id:      (*uuid.UUID)(bank.ID),
		Name:    bank.Name,
		Address: bank.Address,
	}
	return responseBank
}
