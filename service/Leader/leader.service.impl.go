package leader

import (
	request "Expire/data/request/Leader"
	response "Expire/data/response/Leader"
	"Expire/helper"
	"Expire/model"
	repository "Expire/repository/Leader"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type LeaderServiceImpl struct {
	LeaderRepository repository.LeaderRepository
	Validate         *validator.Validate
}

func NewLeaderServiceImpl(leaderRepository repository.LeaderRepository, validate *validator.Validate) LeaderService {
	return &LeaderServiceImpl{
		LeaderRepository: leaderRepository,
		Validate:         validate,
	}
}

func (t LeaderServiceImpl) Create(leader request.CreateLeaderRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(leader)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	leaderModel := model.Leader{
		Name:    leader.Name,
		NIP:     leader.NIP,
		Phone:   leader.Phone,
		Jabatan: leader.Jabatan,
	}
	createError := t.LeaderRepository.Create(leaderModel)

	if createError != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return nil
}

func (t LeaderServiceImpl) GetLeader(id string) (*response.LeaderResponse, *helper.CustomError) {
	result, fetchError := t.LeaderRepository.GetLeader(id)

	if fetchError != nil {
		return nil, fetchError
	} else {
		response := response.LeaderResponse{
			Id:      (*uuid.UUID)(result.ID),
			Name:    result.Name,
			NIP:     result.NIP,
			Phone:   result.Phone,
			Jabatan: result.Jabatan,
		}

		return &response, nil
	}
}

func (t LeaderServiceImpl) GetAllLeader() ([]response.LeaderResponse, *helper.CustomError) {
	result, fetchError := t.LeaderRepository.GetAllLeader()

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapLeadersToLeaderResponse(result), nil
	}
}

func (t LeaderServiceImpl) mapLeadersToLeaderResponse(leaders []model.Leader) []response.LeaderResponse {
	responseLeaders := make([]response.LeaderResponse, len(leaders))
	for i, leader := range leaders {
		responseLeaders[i] = t.convertLeaderToLeaderResponse(leader)
	}
	return responseLeaders
}

func (t LeaderServiceImpl) convertLeaderToLeaderResponse(leader model.Leader) response.LeaderResponse {
	// Perform necessary conversion logic here, potentially selecting specific fields
	responseLeader := response.LeaderResponse{
		Id:      (*uuid.UUID)(leader.ID),
		Name:    leader.Name,
		NIP:     leader.NIP,
		Phone:   leader.Phone,
		Jabatan: leader.Jabatan,
	}
	return responseLeader
}
