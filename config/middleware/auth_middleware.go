package middleware

import (
	"Expire/config"
	authentication "Expire/data/request/Authentication"
	"Expire/helper"
	"Expire/model"
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		env, _ := config.LoadConfig(".")
		tokenString := helper.ExtractToken(c)

		_, err := helper.ValidateTokenFormat(tokenString, env.AccessTokenPublicKey)
		fmt.Println("🐭", err)

		if err != nil {
			fileName, atLine := helper.GetFileAndLine(err)
			helper.ResponseError(c, helper.CustomError{
				Code:     401,
				Message:  "Check you credential.",
				FileName: fileName,
				AtLine:   atLine,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func DeserializeUser(DB *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Write code for Serialize Bearer Token.
		var access_token string
		authorization := ctx.GetHeader("Authorization")
		accessTokenFromCookie, errorCookie := ctx.Cookie("access_token")
		if strings.HasPrefix(authorization, "Bearer ") {
			access_token = strings.TrimPrefix(authorization, "Bearer ")
		} else if accessTokenFromCookie != "" {
			access_token = accessTokenFromCookie
		}

		if access_token == "" {
			fileName, atLine := helper.GetFileAndLine(errorCookie)
			helper.ResponseError(ctx, helper.CustomError{
				Code:     401,
				Message:  "Unauthorize.",
				FileName: fileName,
				AtLine:   atLine,
			})
			ctx.Abort()
			return
		}

		env, _ := config.LoadConfig(".")

		identifier, errTokenClaims := helper.ExtractIdentifierFromToken(access_token, env.AccessTokenPublicKey)

		if errTokenClaims != nil {
			fileName, atLine := helper.GetFileAndLine(errTokenClaims)
			helper.ResponseError(ctx, helper.CustomError{
				Code:     401,
				Message:  "Unauthorize.",
				FileName: fileName,
				AtLine:   atLine,
			})
			ctx.Abort()
		}

		contextTodo := context.TODO()
		userId, errRedis := config.RedisClient.Get(contextTodo, *identifier).Result()

		if errRedis == redis.Nil {
			fileName, atLine := helper.GetFileAndLine(errTokenClaims)
			helper.ResponseError(ctx, helper.CustomError{
				Code:     401,
				Message:  "Unauthorize.",
				FileName: fileName,
				AtLine:   atLine,
			})
			ctx.Abort()
			return
		}

		var user model.User
		errDB := DB.First(&user, "id = ?", userId).Error
		if errDB == gorm.ErrRecordNotFound {
			fileName, atLine := helper.GetFileAndLine(errTokenClaims)
			helper.ResponseError(ctx, helper.CustomError{
				Code:     401,
				Message:  "Unauthorize.",
				FileName: fileName,
				AtLine:   atLine,
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", authentication.FilterUserRecord(&user))
		ctx.Set("access_token_uuid", identifier)

		ctx.Next()
	}
}

/*
 Redis should store the User ID, with Key Token.ID or Token
 So when log out, we can extracting the Token.ID or Token as a Key, and get the User ID value for Query processing
 Task:
 - Change Key to Token ID / Token Instead of UserID, in Login and RefreshToken
 - Open code for Extracting Token ID value from Redis to get User ID in Logout Function
*/
