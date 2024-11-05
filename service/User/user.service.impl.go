package user

import (
	"Expire/helper"
	"Expire/model"
	"strconv"

	request "Expire/data/request/User"
	response "Expire/data/response/User"

	bankRepository "Expire/repository/Bank"
	externalRepository "Expire/repository/External"
	leaderRepository "Expire/repository/Leader"
	supervisorRepository "Expire/repository/Supervisor"
	repository "Expire/repository/User"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository       repository.UserRepository
	SupervisorRepository supervisorRepository.SupervisorRepository
	LeaderRepository     leaderRepository.LeaderRepository
	ExternalRepository   externalRepository.ExternalRepository
	BankRepository       bankRepository.BankRepository
	Validate             *validator.Validate
}

func NewUserServiceImpl(
	bankRepository bankRepository.BankRepository,
	userRepository repository.UserRepository,
	supervisorRepository supervisorRepository.SupervisorRepository,
	externalRepository externalRepository.ExternalRepository,
	leaderRepository leaderRepository.LeaderRepository,
	validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository:       userRepository,
		SupervisorRepository: supervisorRepository,
		LeaderRepository:     leaderRepository,
		ExternalRepository:   externalRepository,
		BankRepository:       bankRepository,
		Validate:             validate,
	}
}

func (t UserServiceImpl) GetUserByID(id string) (*response.UserResponse, *helper.CustomError) {
	result, err := t.UserRepository.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	userResponse := response.UserResponse{
		Id:    result.ID,
		Email: result.Email,
	}

	return &userResponse, nil
}

func (t UserServiceImpl) CreateNewUser(data request.CreateUserRequest) *helper.CustomError {

	if data.Type == 1 {
		println("🐷 x")
		supervisorModel := model.Supervisor{
			Name:    data.Name,
			Jabatan: data.Jabatan,
			NIP:     data.NIP,
			Phone:   data.Phone,
		}
		createError := t.SupervisorRepository.Create(supervisorModel)

		println("🐷 f")
		res, errSup := t.SupervisorRepository.FindByName(data.Name)

		println("🐷 b")
		if createError != nil || errSup != nil {
			println("🐷 b")
			return createError
		}

		userModel := model.User{
			Email:    data.Email,
			Password: data.Password,
			Type:     1,
			Key:      res.ID.String(),
		}

		t.UserRepository.Create(userModel)
		println("🐷 h")

	} else if data.Type == 99 {

		leaderModel := model.Leader{
			Name:    data.Name,
			Jabatan: data.Jabatan,
			NIP:     data.NIP,
			Phone:   data.Phone,
		}

		createError := t.LeaderRepository.Create(leaderModel)

		res, errSup := t.LeaderRepository.FindByName(data.Name)

		if createError != nil || errSup != nil {
			return createError
		}

		userModel := model.User{
			Email:    data.Email,
			Password: data.Password,
			Type:     1,
			Key:      res.ID.String(),
		}

		t.UserRepository.Create(userModel)

	} else if data.Type == 2 {

		// Get Bank ID
		bankResult, errBank := t.BankRepository.FindBankById(data.BankID)
		if errBank != nil {
			return errBank
		}

		externalModel := model.External{
			Name:  data.Name,
			Phone: data.Phone,
			Bank:  *bankResult,
		}

		createError := t.ExternalRepository.Create(externalModel)

		res, errSup := t.ExternalRepository.FindByName(data.Name)

		if createError != nil || errSup != nil {
			return createError
		}

		userModel := model.User{
			Email:    data.Email,
			Password: data.Password,
			Type:     2,
			Key:      res.ID.String(),
		}

		t.UserRepository.Create(userModel)
	}

	return nil
}

func (t UserServiceImpl) GetAllUser() ([]response.UserDataResponse, *helper.CustomError) {
	result, fetchError := t.UserRepository.GetAllUser()

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapUsertoUserResponse(result), nil
	}
}

func (t UserServiceImpl) mapUsertoUserResponse(users []model.User) []response.UserDataResponse {
	responseUser := make([]response.UserDataResponse, len(users))
	for i, user := range users {
		responseUser[i] = t.convertUserToUserResponse(user)
	}
	return responseUser
}

func (t UserServiceImpl) convertUserToUserResponse(user model.User) response.UserDataResponse {
	// Perform necessary conversion logic here, potentially selecting specific fields
	responseUser := response.UserDataResponse{
		Id:       user.ID,
		Name:     user.Email,
		Email:    user.Email,
		Phone:    user.ID.String(),
		Type:     strconv.Itoa(user.Type),
		Jabatan:  user.Email,
		NIP:      user.Email,
		BankName: user.Key,
	}
	return responseUser
}
