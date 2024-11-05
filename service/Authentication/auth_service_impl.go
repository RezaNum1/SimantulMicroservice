package authentication

import (
	"Expire/config"
	authentication "Expire/data/request/Authentication"
	"Expire/helper"
	"Expire/model"
	externalRepository "Expire/repository/External"
	leaderRepository "Expire/repository/Leader"
	supervisorRepository "Expire/repository/Supervisor"
	userRepository "Expire/repository/User"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	UserRepository       userRepository.UserRepository
	SupervisorRepository supervisorRepository.SupervisorRepository
	LeaderRepository     leaderRepository.LeaderRepository
	ExternalRepository   externalRepository.ExternalRepository
	Validate             *validator.Validate
}

func NewAuthServiceImpl(userRepository userRepository.UserRepository, supervisorRepository supervisorRepository.SupervisorRepository, leaderRepository leaderRepository.LeaderRepository, externalRepository externalRepository.ExternalRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository:       userRepository,
		SupervisorRepository: supervisorRepository,
		LeaderRepository:     leaderRepository,
		ExternalRepository:   externalRepository,
		Validate:             validate,
	}
}

func (t AuthServiceImpl) Register(user model.User) *helper.CustomError {
	errStructure := t.Validate.Struct(user)

	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	createError := t.UserRepository.Create(user)

	if createError != nil {
		fmt.Println("🐻", createError)
		return createError
	}

	return nil
}

func (t AuthServiceImpl) Login(payload authentication.SignInInput) (*config.TokenDetails, *config.TokenDetails, *model.User, string, *helper.CustomError) {
	user, err := t.UserRepository.ValidateUserAccount(payload)

	name := ""
	if user.Type == 1 {
		supervisor, supErr := t.SupervisorRepository.GetSupervisor(user.Key)
		if supErr != nil {
			return nil, nil, nil, name, supErr
		}
		name = supervisor.Name
	} else if user.Type == 2 {
		external, externalErr := t.ExternalRepository.GetExternal(user.Key)
		if externalErr != nil {
			return nil, nil, nil, name, externalErr
		}
		name = external.Name
	} else if user.Type == 99 {
		leader, leadErr := t.LeaderRepository.GetLeader(user.Key)
		if leadErr != nil {
			return nil, nil, nil, name, leadErr
		}
		name = leader.Name
	} else if user.Type == 3 {
		name = "Administrator"
	}

	if err != nil {
		return nil, nil, nil, name, err
	}

	env, _ := config.LoadConfig(".")

	accessTokenDetails, accessTokenErr := config.CreateToken(user, env.AccessTokenExpiresIn, env.AccessTokenPrivateKey)
	if accessTokenErr != nil {
		fileName, atLine := helper.GetFileAndLine(accessTokenErr)
		return nil, nil, nil, name, &helper.CustomError{
			Code:     400,
			Message:  "Error creating access token.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	refreshTokenDetails, refreshTokenErr := config.CreateToken(user, env.RefreshTokenExpiresIn, env.RefreshTokenPrivateKey)
	if refreshTokenErr != nil {
		fileName, atLine := helper.GetFileAndLine(accessTokenErr)
		return nil, nil, nil, name, &helper.CustomError{
			Code:     400,
			Message:  "Error creating refresh token.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return accessTokenDetails, refreshTokenDetails, user, name, nil
}

func (t AuthServiceImpl) CheckRegisteredEmail(payload authentication.VerifyForgetPassword) bool {
	user, _ := t.UserRepository.FindUserByEmail(payload.Email)

	if user != nil {
		return true
	} else {
		return false
	}
}

func (t AuthServiceImpl) ResetPassword(payload authentication.ResetPassword) *helper.CustomError {
	err := t.UserRepository.UpdatePasssword(payload.Email, payload.NewPassword)

	if err != nil {
		return err
	} else {
		return nil
	}
}
