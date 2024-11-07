package report

import (
	"Expire/helper"
	"Expire/model"
)

type ReportRepository interface {
	Create(report model.Report) *helper.CustomError
	GetReport(id string) (*model.Report, *helper.CustomError)
	GetAllReport() ([]model.Report, *helper.CustomError)
	Update(report model.Report) *helper.CustomError
	DeleteBySupervisor(id string) *helper.CustomError
	DeleteByLeader(id string) *helper.CustomError
	GetAllSupervisorReports(id string) ([]model.Report, *helper.CustomError)
	GetAllLeaderReports(id string) ([]model.Report, *helper.CustomError)
	GetAllBankReports(id string) ([]model.Report, *helper.CustomError)
}
