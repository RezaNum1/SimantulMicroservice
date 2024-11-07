package report

import (
	request "Expire/data/request/Report"
	response "Expire/data/response/Report"
	"Expire/helper"
)

type ReportService interface {
	Create(report request.CreateReportRequest) *helper.CustomError
	GetReport(id string) (*response.ReportResponse, *helper.CustomError)
	GetAllReport() ([]response.ReportResponse, *helper.CustomError)
	GetAllSupervisorReports(id string) ([]response.ReportResponse, *helper.CustomError)
	GetAllLeaderReports(id string) ([]response.ReportResponse, *helper.CustomError)
	GetAllBankReports(id string) ([]response.ReportResponse, *helper.CustomError)
	Update(report request.UpdateReportRequest) *helper.CustomError
	UpdateStatus(status request.UpdateStatusReportRequest) *helper.CustomError
	UpdateBankReport(report request.UpdateBankReportRequest) *helper.CustomError
	UpdateDokumenTemuanStatus(report request.UpdateDokumenTemuanReportRequest) *helper.CustomError
	UpdateDokumenTindakLanjutStatus(report request.UpdateDokumenTindakLanjutReportRequest) *helper.CustomError
}
