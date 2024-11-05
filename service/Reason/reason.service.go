package reason

import (
	request "Expire/data/request/Reason"
	response "Expire/data/response/Reason"
	"Expire/helper"
)

type ReasonService interface {
	Create(report request.CreateReasonRequest) *helper.CustomError
	GetReason(id string) (*response.ReasonResponse, *helper.CustomError)
	GetAllReason() ([]response.ReasonResponse, *helper.CustomError)
	FindReasonsByReportID(reportId string) ([]response.ReasonResponse, *helper.CustomError)
}
