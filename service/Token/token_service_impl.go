package token

import (
	"Expire/config"
	"Expire/helper"
	repository "Expire/repository/User"
	"fmt"
)

type TokenServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewTokenServiceImpl(userRepository repository.UserRepository) TokenService {
	return &TokenServiceImpl{
		UserRepository: userRepository,
	}
}

func (t TokenServiceImpl) RefreshAccessToken(userId string) (*config.TokenDetails, *config.TokenDetails, *helper.CustomError) {
	user, err := t.UserRepository.FindUserById(userId)
	if err != nil {
		fmt.Println("Debug 4")
		return nil, nil, err
	}

	env, _ := config.LoadConfig(".")

	accessTokenDetails, accessTokenErr := config.CreateToken(user, env.AccessTokenExpiresIn, env.AccessTokenPrivateKey)
	if accessTokenErr != nil {
		fileName, atLine := helper.GetFileAndLine(accessTokenErr)
		return nil, nil, &helper.CustomError{
			Code:     400,
			Message:  "Error Creating Access Token.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	refreshTokenDetails, refreshTokenErr := config.CreateToken(user, env.RefreshTokenExpiresIn, env.RefreshTokenPrivateKey)
	if refreshTokenErr != nil {
		fileName, atLine := helper.GetFileAndLine(accessTokenErr)
		return nil, nil, &helper.CustomError{
			Code:     400,
			Message:  "Error creating refresh token.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return accessTokenDetails, refreshTokenDetails, nil
}
