package authentication

import (
	"Expire/config"
	authentication "Expire/data/request/Authentication"
	"Expire/helper"
	"Expire/model"
)

type AuthService interface {
	Register(user model.User) *helper.CustomError
	Login(payload authentication.SignInInput) (*config.TokenDetails, *config.TokenDetails, *model.User, string, *helper.CustomError)
	CheckRegisteredEmail(payload authentication.VerifyForgetPassword) bool
	ResetPassword(payload authentication.ResetPassword) *helper.CustomError
}
