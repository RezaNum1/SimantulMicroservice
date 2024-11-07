package controller

import (
	request "Expire/data/request/Reason"
	response "Expire/data/response"
	"Expire/helper"
	service "Expire/service/Reason"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReasonController struct {
	reasonService service.ReasonService
}

func NewReasonController(service service.ReasonService) *ReasonController {
	return &ReasonController{reasonService: service}
}

func (controller *ReasonController) Create(ctx *gin.Context) {

	var payload request.CreateReasonRequest

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

	err := controller.reasonService.Create(payload)

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

func (controller *ReasonController) GetReason(ctx *gin.Context) {
	stringId := ctx.Param("id")

	reasonResponse, errReasonResponse := controller.reasonService.GetReason(stringId)

	if errReasonResponse != nil {
		helper.ResponseError(ctx, *errReasonResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    reasonResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReasonController) GetAllReason(ctx *gin.Context) {
	reasonResponse, errReasonResponse := controller.reasonService.GetAllReason()

	if errReasonResponse != nil {
		helper.ResponseError(ctx, *errReasonResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    reasonResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ReasonController) FindReasonsByReportID(ctx *gin.Context) {
	stringId := ctx.Param("id")

	reasonResponse, errReasonResponse := controller.reasonService.FindReasonsByReportID(stringId)

	if errReasonResponse != nil {
		helper.ResponseError(ctx, *errReasonResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    reasonResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}
