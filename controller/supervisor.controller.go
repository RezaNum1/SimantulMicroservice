package controller

import (
	request "Expire/data/request/Supervisor"
	response "Expire/data/response"
	"Expire/helper"
	service "Expire/service/Supervisor"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SupervisorController struct {
	supervisorService service.SupervisorService
}

func NewSupervisorController(service service.SupervisorService) *SupervisorController {
	return &SupervisorController{supervisorService: service}
}

func (controller *SupervisorController) Create(ctx *gin.Context) {

	var payload request.CreateSupervisorRequest

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

	err := controller.supervisorService.Create(payload)

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

func (controller *SupervisorController) GetSupervisor(ctx *gin.Context) {
	stringId := ctx.Param("id")

	supervisorResponse, errSupervisorResponse := controller.supervisorService.GetSupervisor(stringId)

	if errSupervisorResponse != nil {
		helper.ResponseError(ctx, *errSupervisorResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    supervisorResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *SupervisorController) GetAllSupervisor(ctx *gin.Context) {
	supervisorResponse, errSupervisorResponse := controller.supervisorService.GetAllSupervisor()

	if errSupervisorResponse != nil {
		helper.ResponseError(ctx, *errSupervisorResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    supervisorResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}
