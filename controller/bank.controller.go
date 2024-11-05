package controller

import (
	request "Expire/data/request/Bank"
	response "Expire/data/response"
	"Expire/helper"
	service "Expire/service/Bank"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	bankService service.BankService
}

func NewBankController(service service.BankService) *BankController {
	return &BankController{bankService: service}
}

func (controller *BankController) Create(ctx *gin.Context) {

	var payload request.CreateBankRequest

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

	err := controller.bankService.Create(payload)

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

func (controller *BankController) GetBank(ctx *gin.Context) {
	stringId := ctx.Param("id")

	bankResponse, errBankResponse := controller.bankService.GetBank(stringId)

	if errBankResponse != nil {
		helper.ResponseError(ctx, *errBankResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    bankResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *BankController) GetAllBank(ctx *gin.Context) {
	bankResponse, errBankResponse := controller.bankService.GetAllBank()

	if errBankResponse != nil {
		helper.ResponseError(ctx, *errBankResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    bankResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}
