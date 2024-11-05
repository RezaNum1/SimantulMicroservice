package controller

import (
	"Expire/config"
	authentication "Expire/data/request/Authentication"
	response "Expire/data/response"
	userResponse "Expire/data/response/User"
	"Expire/helper"
	"Expire/model"
	service "Expire/service/Authentication"
	"context"
	"time"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{authService: service}
}

func (controller *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *authentication.SignUpInput

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

	errors := helper.ValidateStruct(payload)

	if errors != nil {
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: "Auth Controller",
			AtLine:   46,
		})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		fileName, atLine := helper.GetFileAndLine(errBindJSON)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     403,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	hashedPassword, errBcrypt := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if errBcrypt != nil {
		fileName, atLine := helper.GetFileAndLine(errBindJSON)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     403,
			Message:  "Failed to Hashed Password.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	newUser := model.User{
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
		Key:      payload.Key,
		Type:     payload.Type,
	}

	errResponse := controller.authService.Register(newUser)

	if errResponse != nil {
		helper.ResponseError(ctx, *errResponse)
	}

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Code:    200,
		Message: "Success",
		Data:    nil,
	})
}

func (controller *AuthController) SignInUser(ctx *gin.Context) {
	var payload *authentication.SignInInput
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

	accessTokenDetails, refreshTokenDetails, user, name, err := controller.authService.Login(*payload)

	if err != nil {
		helper.ResponseError(ctx, *err)
		return
	}

	contextID := context.TODO()
	now := time.Now()

	errAccess := config.RedisClient.Set(contextID, accessTokenDetails.Identifier, accessTokenDetails.UserID, time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		fileName, atLine := helper.GetFileAndLine(errAccess)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     403,
			Message:  "Username or Password is Wrong.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	errRefresh := config.RedisClient.Set(contextID, refreshTokenDetails.Identifier, refreshTokenDetails.UserID, time.Unix(*refreshTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errRefresh != nil {
		fileName, atLine := helper.GetFileAndLine(errRefresh)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     403,
			Message:  "Username or Password is Wrong.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	// Set Cookie Here with Gin
	ctx.SetCookie("access_token", *accessTokenDetails.Token, 3600, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", *refreshTokenDetails.Token, 3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Code:    200,
		Message: "Success",
		Data: userResponse.LoginResponse{
			AccessToken:  *accessTokenDetails.Token,
			RefreshToken: *refreshTokenDetails.Token,
			Type:         user.Type,
			Id:           user.Key,
			Name:         name,
		},
	})
}

func (controller *AuthController) Logout(ctx *gin.Context) {
	refresh_token, errCookie := ctx.Cookie("refresh_token")

	if errCookie != nil {
		fileName, atLine := helper.GetFileAndLine(errCookie)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	env, _ := config.LoadConfig(".")
	identifier, errValidateToken := helper.ExtractIdentifierFromToken(refresh_token, env.RefreshTokenPublicKey) // you will get UserID & TokenUUID
	if errValidateToken != nil {
		fileName, atLine := helper.GetFileAndLine(errValidateToken)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     401,
			Message:  "Invalid Token.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	contextID := context.TODO()

	access_token_uuid, _ := ctx.Get("access_token_uuid")
	_, errDelRedis := config.RedisClient.Del(contextID, *identifier, access_token_uuid.(string)).Result()

	if errDelRedis != nil {
		fileName, atLine := helper.GetFileAndLine(errValidateToken)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     401,
			Message:  "Failed Delete Redis Property.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	// Set Cookie Here with Gin
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Code:    200,
		Message: "Success",
		Data:    nil,
	})
}

func (controller *AuthController) VerifyForgetPassword(ctx *gin.Context) {
	var payload *authentication.VerifyForgetPassword

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

	errors := helper.ValidateStruct(payload)

	if errors != nil {
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: "Auth Controller",
			AtLine:   241,
		})
		return
	}

	isEmailRegistered := controller.authService.CheckRegisteredEmail(*payload)

	if isEmailRegistered {
		// Send Verification Email Here
		ctx.JSON(http.StatusOK, response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data: userResponse.VerifyForgetPassword{
				Registered: true,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data: userResponse.VerifyForgetPassword{
				Registered: false,
			},
		})
	}

}

func (controller *AuthController) ResetPassword(ctx *gin.Context) {
	var payload *authentication.ResetPassword

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

	errors := helper.ValidateStruct(payload)

	if errors != nil {
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: "Auth Controller",
			AtLine:   241,
		})
		return
	}

	errorDB := controller.authService.ResetPassword(*payload)

	if errorDB != nil {
		helper.ResponseError(ctx, *errorDB)
	} else {
		ctx.JSON(http.StatusOK, response.Response{
			Success: true,
			Code:    200,
			Message: "Success",
			Data: userResponse.ResetPassword{
				PasswordValid: true,
			},
		})
	}

}
