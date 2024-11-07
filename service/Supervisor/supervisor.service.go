package supervisor

import (
	request "Expire/data/request/Supervisor"
	response "Expire/data/response/Supervisor"
	"Expire/helper"
)

type SupervisorService interface {
	Create(report request.CreateSupervisorRequest) *helper.CustomError
	GetSupervisor(id string) (*response.SupervisorResponse, *helper.CustomError)
	GetAllSupervisor() ([]response.SupervisorResponse, *helper.CustomError)
	Delete(id string) *helper.CustomError
}
