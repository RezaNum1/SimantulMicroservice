package token

import (
	"Expire/config"
	"Expire/helper"
)

type TokenService interface {
	RefreshAccessToken(userId string) (*config.TokenDetails, *config.TokenDetails, *helper.CustomError)
}
