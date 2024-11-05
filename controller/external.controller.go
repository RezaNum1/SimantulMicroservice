package controller

import (
	request "Expire/data/request/External"
	response "Expire/data/response"
	"Expire/helper"
	service "Expire/service/External"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExternalController struct {
	externalService service.ExternalService
}

func NewExternalController(service service.ExternalService) *ExternalController {
	return &ExternalController{externalService: service}
}

func (controller *ExternalController) Create(ctx *gin.Context) {

	var payload request.CreateExternalRequest

	errBindJSON := ctx.ShouldBindJSON(&payload)
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

	err := controller.externalService.Create(payload)

	if err != nil {
		println("🐶 1")
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

func (controller *ExternalController) GetExternal(ctx *gin.Context) {
	stringId := ctx.Param("id")

	externalResponse, errExternalResponse := controller.externalService.GetExternal(stringId)

	if errExternalResponse != nil {
		helper.ResponseError(ctx, *errExternalResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    externalResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ExternalController) GetAllExternal(ctx *gin.Context) {
	externalResponse, errExternalResponse := controller.externalService.GetAllExternal()

	if errExternalResponse != nil {
		helper.ResponseError(ctx, *errExternalResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    externalResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *ExternalController) Update(ctx *gin.Context) {

	var payload request.UpdateExternalRequest

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

	err := controller.externalService.Update(payload)

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
