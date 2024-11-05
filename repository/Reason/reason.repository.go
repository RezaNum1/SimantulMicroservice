package reason

import (
	"Expire/helper"
	"Expire/model"
)

type ReasonRepository interface {
	Create(report model.Reason) *helper.CustomError
	GetReason(id string) (*model.Reason, *helper.CustomError)
	GetAllReason() ([]model.Reason, *helper.CustomError)
	FindReasonsByReportID(reportId string) ([]model.Reason, *helper.CustomError)
}
