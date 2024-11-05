package supervisor

import (
	"Expire/helper"
	"Expire/model"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupervisorRepositoryImpl struct {
	Db *gorm.DB
}

func NewSupervisorRepositoryImpl(Db *gorm.DB) SupervisorRepository {
	return &SupervisorRepositoryImpl{Db: Db}
}

func (t *SupervisorRepositoryImpl) Create(report model.Supervisor) *helper.CustomError {
	result := t.Db.Create(&report)

	if result.Error != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Creating New Report.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return nil
}

func (t *SupervisorRepositoryImpl) GetSupervisor(id string) (*model.Supervisor, *helper.CustomError) {
	var supervisor model.Supervisor

	supervisorId, err := uuid.Parse(id)
	result := t.Db.First(&supervisor, "id = ?", supervisorId)
	if result.Error != nil || err != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Supervisor",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return &supervisor, nil
}

func (t *SupervisorRepositoryImpl) GetAllSupervisor() ([]model.Supervisor, *helper.CustomError) {
	var supervisors []model.Supervisor
	result := t.Db.Find(&supervisors)
	if result.Error != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reports",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return supervisors, nil
}

func (t *SupervisorRepositoryImpl) FindSupervisorById(supervisorId string) (*model.Supervisor, *helper.CustomError) {
	var err error
	supervisor := model.Supervisor{}

	if err = t.Db.Model(model.Supervisor{}).First(&supervisor, "id = ?", strings.ToLower(supervisorId)).Error; err != nil {
		return nil, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}
	return &supervisor, nil
}

func (t *SupervisorRepositoryImpl) FindByName(name string) (*model.Supervisor, *helper.CustomError) {
	var err error
	supervisor := model.Supervisor{}

	if err = t.Db.Model(model.Supervisor{}).First(&supervisor, "name = ?", name).Error; err != nil {
		return nil, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}
	return &supervisor, nil
}
