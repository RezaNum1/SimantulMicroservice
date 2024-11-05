package supervisor

import (
	"Expire/helper"
	"Expire/model"
)

type SupervisorRepository interface {
	Create(report model.Supervisor) *helper.CustomError
	GetSupervisor(id string) (*model.Supervisor, *helper.CustomError)
	GetAllSupervisor() ([]model.Supervisor, *helper.CustomError)
	FindSupervisorById(supervisorId string) (*model.Supervisor, *helper.CustomError)
	FindByName(name string) (*model.Supervisor, *helper.CustomError)
}
