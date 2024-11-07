package controller

import (
	request "Expire/data/request/Report"
	response "Expire/data/response"
	"Expire/helper"
	service "Expire/service/Report"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportService service.ReportService
}

func NewReportController(service service.ReportService) *ReportController {
	return &ReportController{reportService: service}
}

func (controller *ReportController) Create(ctx *gin.Context) {

	var payload request.CreateReportRequest
	println("🐶 x")
	errBindJSON := ctx.ShouldBindJSON(&payload)
	println("🐶 1")
	if errBindJSON != nil {
		println("🐶 2")
		fileName, atLine := helper.GetFileAndLine(errBindJSON)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	err := controller.reportService.Create(payload)

	if err != nil {
		helper.ResponseError(ctx, *err)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    nil,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReportController) GetReport(ctx *gin.Context) {
	stringId := ctx.Param("id")

	reportResponse, errReportResponse := controller.reportService.GetReport(stringId)

	if errReportResponse != nil {
		helper.ResponseError(ctx, *errReportResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    reportResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReportController) GetAllReport(ctx *gin.Context) {
	reportResponse, errReportResponse := controller.reportService.GetAllReport()

	if errReportResponse != nil {
		helper.ResponseError(ctx, *errReportResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    reportResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReportController) GetAllSupervisorReport(ctx *gin.Context) {
	stringId := ctx.Param("id")
	reportResponse, errReportResponse := controller.reportService.GetAllSupervisorReports(stringId)

	if errReportResponse != nil {
		helper.ResponseError(ctx, *errReportResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    reportResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReportController) GetAllLeaderReport(ctx *gin.Context) {
	stringId := ctx.Param("id")
	reportResponse, errReportResponse := controller.reportService.GetAllLeaderReports(stringId)

	if errReportResponse != nil {
		helper.ResponseError(ctx, *errReportResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    reportResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReportController) GetAllBankReport(ctx *gin.Context) {
	stringId := ctx.Param("id")
	reportResponse, errReportResponse := controller.reportService.GetAllBankReports(stringId)

	if errReportResponse != nil {
		helper.ResponseError(ctx, *errReportResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    reportResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReportController) Update(ctx *gin.Context) {

	var payload request.UpdateReportRequest

	errBindJSON := ctx.ShouldBindJSON(&payload)
	if errBindJSON != nil {
		fileName, atLine := helper.GetFileAndLine(errBindJSON)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	err := controller.reportService.Update(payload)

	if err != nil {
		helper.ResponseError(ctx, *err)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    nil,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReportController) UpdateStatus(ctx *gin.Context) {

	var payload request.UpdateStatusReportRequest

	errBindJSON := ctx.ShouldBindJSON(&payload)
	if errBindJSON != nil {
		println("🐶 4")
		fileName, atLine := helper.GetFileAndLine(errBindJSON)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	err := controller.reportService.UpdateStatus(payload)

	if err != nil {
		helper.ResponseError(ctx, *err)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    nil,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReportController) UpdateBankReport(ctx *gin.Context) {

	var payload request.UpdateBankReportRequest

	errBindJSON := ctx.ShouldBindJSON(&payload)
	if errBindJSON != nil {
		println("🐶 1")
		fileName, atLine := helper.GetFileAndLine(errBindJSON)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	err := controller.reportService.UpdateBankReport(payload)

	if err != nil {
		helper.ResponseError(ctx, *err)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    nil,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReportController) UpdateDokumenTemuan(ctx *gin.Context) {

	var payload request.UpdateDokumenTemuanReportRequest

	errBindJSON := ctx.ShouldBindJSON(&payload)
	if errBindJSON != nil {
		println("🐶 4")
		fileName, atLine := helper.GetFileAndLine(errBindJSON)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	err := controller.reportService.UpdateDokumenTemuanStatus(payload)

	if err != nil {
		helper.ResponseError(ctx, *err)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    nil,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReportController) UpdateDokumenTindakLanjut(ctx *gin.Context) {

	var payload request.UpdateDokumenTindakLanjutReportRequest

	errBindJSON := ctx.ShouldBindJSON(&payload)
	if errBindJSON != nil {
		println("🐶 4")
		fileName, atLine := helper.GetFileAndLine(errBindJSON)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	err := controller.reportService.UpdateDokumenTindakLanjutStatus(payload)

	if err != nil {
		helper.ResponseError(ctx, *err)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    nil,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}
