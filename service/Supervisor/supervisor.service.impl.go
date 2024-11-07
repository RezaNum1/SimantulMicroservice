package supervisor

import (
	request "Expire/data/request/Supervisor"
	response "Expire/data/response/Supervisor"
	"Expire/helper"
	"Expire/model"
	repository "Expire/repository/Supervisor"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SupervisorServiceImpl struct {
	SupervisorRepository repository.SupervisorRepository
	Validate             *validator.Validate
}

func NewSupervisorServiceImpl(supervisorRepository repository.SupervisorRepository, validate *validator.Validate) SupervisorService {
	return &SupervisorServiceImpl{
		SupervisorRepository: supervisorRepository,
		Validate:             validate,
	}
}

func (t SupervisorServiceImpl) Create(supervisor request.CreateSupervisorRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(supervisor)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	supervisorModel := model.Supervisor{
		Name:    supervisor.Name,
		Jabatan: supervisor.Jabatan,
		NIP:     supervisor.NIP,
		Phone:   supervisor.Phone,
	}
	createError := t.SupervisorRepository.Create(supervisorModel)

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

func (t SupervisorServiceImpl) GetSupervisor(id string) (*response.SupervisorResponse, *helper.CustomError) {
	result, fetchError := t.SupervisorRepository.GetSupervisor(id)

	if fetchError != nil {
		return nil, fetchError
	} else {
		response := response.SupervisorResponse{
			Id:      (*uuid.UUID)(result.ID),
			Name:    result.Name,
			Jabatan: result.Jabatan,
			Phone:   result.Phone,
			NIP:     result.NIP,
		}

		return &response, nil
	}
}

func (t SupervisorServiceImpl) GetAllSupervisor() ([]response.SupervisorResponse, *helper.CustomError) {
	result, fetchError := t.SupervisorRepository.GetAllSupervisor()

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapSupervisorsToSupervisorResponse(result), nil
	}
}

func (t SupervisorServiceImpl) mapSupervisorsToSupervisorResponse(supervisors []model.Supervisor) []response.SupervisorResponse {
	responseSupervisors := make([]response.SupervisorResponse, len(supervisors))
	for i, supervisor := range supervisors {
		responseSupervisors[i] = t.convertSupervisorToSupervisorResponse(supervisor)
	}
	return responseSupervisors
}

func (t SupervisorServiceImpl) convertSupervisorToSupervisorResponse(supervisor model.Supervisor) response.SupervisorResponse {
	// Perform necessary conversion logic here, potentially selecting specific fields
	responseSupervisor := response.SupervisorResponse{
		Id:      (*uuid.UUID)(supervisor.ID),
		Name:    supervisor.Name,
		Jabatan: supervisor.Jabatan,
		Phone:   supervisor.Phone,
		NIP:     supervisor.NIP,
	}
	return responseSupervisor
}

func (t SupervisorServiceImpl) FindByName(id string) (*response.SupervisorResponse, *helper.CustomError) {
	result, fetchError := t.SupervisorRepository.GetSupervisor(id)

	if fetchError != nil {
		return nil, fetchError
	} else {
		response := response.SupervisorResponse{
			Id:      (*uuid.UUID)(result.ID),
			Name:    result.Name,
			Jabatan: result.Jabatan,
			Phone:   result.Phone,
			NIP:     result.NIP,
		}

		return &response, nil
	}
}

func (t SupervisorServiceImpl) Delete(id string) *helper.CustomError {
	err := t.SupervisorRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
