package external

import (
	request "Expire/data/request/External"
	response "Expire/data/response/External"
	"Expire/helper"
)

type ExternalService interface {
	Create(report request.CreateExternalRequest) *helper.CustomError
	GetExternal(id string) (*response.ExternalResponse, *helper.CustomError)
	GetAllExternal() ([]response.ExternalResponse, *helper.CustomError)
	Update(external request.UpdateExternalRequest) *helper.CustomError
}
