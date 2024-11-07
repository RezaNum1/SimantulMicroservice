package reason

import (
	request "Expire/data/request/Reason"
	response "Expire/data/response/Reason"
	"Expire/helper"
	"Expire/model"
	repository "Expire/repository/Reason"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ReasonServiceImpl struct {
	ReasonRepository repository.ReasonRepository
	Validate         *validator.Validate
}

func NewReasonServiceImpl(reasonRepository repository.ReasonRepository, validate *validator.Validate) ReasonService {
	return &ReasonServiceImpl{
		ReasonRepository: reasonRepository,
		Validate:         validate,
	}
}

func (t ReasonServiceImpl) Create(reason request.CreateReasonRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(reason)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	reasonModel := model.Reason{
		Description: reason.Description,
		ReportID:    reason.ReportID,
	}
	createError := t.ReasonRepository.Create(reasonModel)

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

func (t ReasonServiceImpl) GetReason(id string) (*response.ReasonResponse, *helper.CustomError) {
	result, fetchError := t.ReasonRepository.GetReason(id)

	if fetchError != nil {
		return nil, fetchError
	} else {
		response := response.ReasonResponse{
			Id:           (*uuid.UUID)(result.ID),
			Description:  result.Description,
			RejectedStep: result.RejectedStep,
		}

		return &response, nil
	}
}

func (t ReasonServiceImpl) GetAllReason() ([]response.ReasonResponse, *helper.CustomError) {
	result, fetchError := t.ReasonRepository.GetAllReason()

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapReasonsToReasonResponse(result), nil
	}
}

func (t ReasonServiceImpl) mapReasonsToReasonResponse(reasons []model.Reason) []response.ReasonResponse {
	responseReasons := make([]response.ReasonResponse, len(reasons))
	for i, reason := range reasons {
		responseReasons[i] = t.convertReasonToReasonResponse(reason)
	}
	return responseReasons
}

func (t ReasonServiceImpl) convertReasonToReasonResponse(reason model.Reason) response.ReasonResponse {
	// Perform necessary conversion logic here, potentially selecting specific fields
	responseReason := response.ReasonResponse{
		Id:           (*uuid.UUID)(reason.ID),
		Description:  reason.Description,
		RejectedStep: reason.RejectedStep,
	}
	return responseReason
}

func (t ReasonServiceImpl) FindReasonsByReportID(reportId string) ([]response.ReasonResponse, *helper.CustomError) {
	result, fetchError := t.ReasonRepository.FindReasonsByReportID(reportId)

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapReasonsToReasonResponse(result), nil
	}
}
