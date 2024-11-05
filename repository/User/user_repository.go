package user

import (
	authentication "Expire/data/request/Authentication"
	"Expire/helper"
	"Expire/model"
)

type UserRepository interface {
	GetUserByID(id string) (model.User, *helper.CustomError)
	Create(user model.User) *helper.CustomError
	ValidateUserAccount(payload authentication.SignInInput) (*model.User, *helper.CustomError)
	FindUserById(userId string) (*model.User, *helper.CustomError)
	FindUserByEmail(email string) (*model.User, *helper.CustomError)
	UpdatePasssword(email string, newPassword string) *helper.CustomError
	GetAllUser() ([]model.User, *helper.CustomError)
}
