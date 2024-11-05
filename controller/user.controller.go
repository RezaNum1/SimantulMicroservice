package controller

import (
	request "Expire/data/request/User"
	response "Expire/data/response"
	"Expire/helper"
	service "Expire/service/User"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{userService: service}
}

func (controller *UserController) GetUser(ctx *gin.Context) {
	// GET USER ID
	userId, err := helper.GetUserId(ctx)
	if err != nil {
		helper.ResponseError(ctx, helper.CustomError{
			Message: "Invalid Identifier",
			Code:    404,
		})
	}

	userResponse, errResponse := controller.userService.GetUserByID(*userId)

	if errResponse != nil {
		helper.ResponseError(ctx, *errResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    userResponse,
		}
		ctx.JSON(http.StatusOK, webResponse)
	}
}

func (controller *UserController) Create(ctx *gin.Context) {
	println("🐷 2")
	var payload request.CreateUserRequest

	errBindJSON := ctx.ShouldBindJSON(&payload)
	println("🐷 5")
	if errBindJSON != nil {
		println("🐷 3")
		fileName, atLine := helper.GetFileAndLine(errBindJSON)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}
	println("🐷 8")
	err := controller.userService.CreateNewUser(payload)

	println("🐷 6")
	if err != nil {
		println("🐷 1")
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

func (controller *UserController) GetAllUser(ctx *gin.Context) {

	userResponse, errResponse := controller.userService.GetAllUser()

	if errResponse != nil {
		helper.ResponseError(ctx, *errResponse)
	} else {
		webResponse := response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data:    userResponse,
		}
		ctx.JSON(http.StatusOK, webResponse)
	}
}
