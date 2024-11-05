package user

import (
	request "Expire/data/request/User"
	response "Expire/data/response/User"
	"Expire/helper"
)

type UserService interface {
	GetUserByID(id string) (*response.UserResponse, *helper.CustomError)
	CreateNewUser(data request.CreateUserRequest) *helper.CustomError
	GetAllUser() ([]response.UserDataResponse, *helper.CustomError)
}
