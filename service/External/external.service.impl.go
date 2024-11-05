package external

import (
	request "Expire/data/request/External"
	response "Expire/data/response/External"
	"Expire/helper"
	"Expire/model"
	bankRepository "Expire/repository/Bank"
	repository "Expire/repository/External"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ExternalServiceImpl struct {
	ExternalRepository repository.ExternalRepository
	BankRepository     bankRepository.BankRepository
	Validate           *validator.Validate
}

func NewExternalServiceImpl(externalRepository repository.ExternalRepository, bankRepository bankRepository.BankRepository, validate *validator.Validate) ExternalService {
	return &ExternalServiceImpl{
		ExternalRepository: externalRepository,
		BankRepository:     bankRepository,
		Validate:           validate,
	}
}

func (t ExternalServiceImpl) Create(external request.CreateExternalRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(external)
	if errStructure != nil {
		println("🐶 2")
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	bankResult, err := t.BankRepository.FindBankById(external.BankID)
	if err != nil {
		println("🐶 3")
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	externalModel := model.External{
		Name:  external.Name,
		Phone: external.Phone,
		Bank:  *bankResult,
	}
	createError := t.ExternalRepository.Create(externalModel)

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

func (t ExternalServiceImpl) GetExternal(id string) (*response.ExternalResponse, *helper.CustomError) {
	result, fetchError := t.ExternalRepository.GetExternal(id)

	if fetchError != nil {
		return nil, fetchError
	} else {
		response := response.ExternalResponse{
			Id:    (*uuid.UUID)(result.ID),
			Name:  result.Name,
			Phone: result.Phone,
			Bank:  result.Bank,
		}

		return &response, nil
	}
}

func (t ExternalServiceImpl) GetAllExternal() ([]response.ExternalResponse, *helper.CustomError) {
	result, fetchError := t.ExternalRepository.GetAllExternal()

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapExternalsToExternalResponse(result), nil
	}
}

func (t ExternalServiceImpl) mapExternalsToExternalResponse(externals []model.External) []response.ExternalResponse {
	responseExternals := make([]response.ExternalResponse, len(externals))
	for i, external := range externals {
		responseExternals[i] = t.convertExternalToExternalResponse(external)
	}
	return responseExternals
}

func (t ExternalServiceImpl) convertExternalToExternalResponse(external model.External) response.ExternalResponse {
	// Perform necessary conversion logic here, potentially selecting specific fields
	responseExternal := response.ExternalResponse{
		Id:    (*uuid.UUID)(external.ID),
		Name:  external.Name,
		Phone: external.Phone,
		Bank:  external.Bank,
	}
	return responseExternal
}

func (t ExternalServiceImpl) Update(external request.UpdateExternalRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(external)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	bank, errBank := t.BankRepository.GetBank(external.BankID)

	data, fetchErr := t.ExternalRepository.GetExternal(external.ID)
	data.Name = external.Name
	data.Phone = external.Phone
	data.Bank = *bank

	saveErr := t.ExternalRepository.Update(*data)

	if fetchErr != nil || saveErr != nil || errBank != nil {
		println("🐶 2")
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
