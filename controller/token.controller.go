package controller

import (
	"Expire/config"
	response "Expire/data/response"
	userResponse "Expire/data/response/User"
	"Expire/helper"
	service "Expire/service/Token"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	tokenService service.TokenService
}

func NewTokenController(tokenService service.TokenService) *TokenController {
	return &TokenController{tokenService: tokenService}
}

func (controller *TokenController) RefreshAccessToken(ctx *gin.Context) {
	var refresh_token string
	authorization := ctx.GetHeader("Authorization")
	if strings.HasPrefix(authorization, "Bearer ") {
		refresh_token = strings.TrimPrefix(authorization, "Bearer ")
	}

	env, _ := config.LoadConfig(".")

	identifier, errorValidation := helper.ExtractIdentifierFromToken(refresh_token, env.RefreshTokenPublicKey) // you will get UserID & TokenUUID
	if errorValidation != nil {
		println("🐱 2")
		fileName, atLine := helper.GetFileAndLine(errorValidation)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     403,
			Message:  "Refresh Token Not Found.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	userId, errTokenClaims := config.RedisClient.Get(ctx, *identifier).Result()

	if errTokenClaims != nil {
		fileName, atLine := helper.GetFileAndLine(errTokenClaims)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     403,
			Message:  "Refresh Token Not Found.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	accessTokenDetails, refreshTokenDetails, errorRefreshAccessToken := controller.tokenService.RefreshAccessToken(userId)
	if errorRefreshAccessToken != nil {
		helper.ResponseError(ctx, *errorRefreshAccessToken)
	}

	contextID := context.TODO()
	now := time.Now()

	errAccess := config.RedisClient.Set(contextID, accessTokenDetails.Identifier, accessTokenDetails.UserID, time.Unix(*accessTokenDetails.ExpiresIn, 0).Sub(now)).Err()
	if errAccess != nil {
		println("🐱 3")
		fileName, atLine := helper.GetFileAndLine(errAccess)
		helper.ResponseError(ctx, helper.CustomError{
			Code:     400,
			Message:  "Bad Request.",
			FileName: fileName,
			AtLine:   atLine,
		})
		return
	}

	ctx.SetCookie("access_token", *accessTokenDetails.Token, 3600, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, response.Response{
		Success: true,
		Code:    200,
		Message: "Success",
		Data: userResponse.TokenResponse{
			AccessToken:  *accessTokenDetails.Token,
			RefreshToken: *refreshTokenDetails.Token,
		},
	})
}
