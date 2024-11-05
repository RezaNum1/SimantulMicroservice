package controller

import (
	request "Expire/data/request/Leader"
	response "Expire/data/response"
	"Expire/helper"
	service "Expire/service/Leader"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LeaderController struct {
	leaderService service.LeaderService
}

func NewLeaderController(service service.LeaderService) *LeaderController {
	return &LeaderController{leaderService: service}
}

func (controller *LeaderController) Create(ctx *gin.Context) {

	var payload request.CreateLeaderRequest

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

	err := controller.leaderService.Create(payload)

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

func (controller *LeaderController) GetLeader(ctx *gin.Context) {
	stringId := ctx.Param("id")

	leaderResponse, errLeaderResponse := controller.leaderService.GetLeader(stringId)

	if errLeaderResponse != nil {
		helper.ResponseError(ctx, *errLeaderResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    leaderResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *LeaderController) GetAllLeader(ctx *gin.Context) {
	leaderResponse, errLeaderResponse := controller.leaderService.GetAllLeader()

	if errLeaderResponse != nil {
		helper.ResponseError(ctx, *errLeaderResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    leaderResponse,
		}

		ctx.JSON(http.StatusOK, webResponse)
	}
}
