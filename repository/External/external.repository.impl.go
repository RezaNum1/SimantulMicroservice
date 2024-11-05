package external

import (
	"Expire/helper"
	"Expire/model"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExternalRepositoryImpl struct {
	Db *gorm.DB
}

func NewExternalRepositoryImpl(Db *gorm.DB) ExternalRepository {
	return &ExternalRepositoryImpl{Db: Db}
}

func (t *ExternalRepositoryImpl) Create(report model.External) *helper.CustomError {
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

func (t *ExternalRepositoryImpl) GetExternal(id string) (*model.External, *helper.CustomError) {
	var external model.External

	externalId, err := uuid.Parse(id)
	result := t.Db.Preload("Bank").First(&external, "id = ?", externalId)
	if result.Error != nil || err != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching External",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return &external, nil
}

func (t *ExternalRepositoryImpl) GetAllExternal() ([]model.External, *helper.CustomError) {
	var externals []model.External
	result := t.Db.Preload("Bank").Find(&externals)
	if result.Error != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reports",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return externals, nil
}

func (t *ExternalRepositoryImpl) FindExternalById(externalId string) (*model.External, *helper.CustomError) {
	var err error
	external := model.External{}

	if err = t.Db.Model(model.External{}).First(&external, "id = ?", strings.ToLower(externalId)).Error; err != nil {
		return nil, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}
	return &external, nil
}

func (t *ExternalRepositoryImpl) Update(external model.External) *helper.CustomError {
	err := t.Db.Save(&external).Error

	if err != nil {
		fileName, atLine := helper.GetFileAndLine(err)
		return &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Creating New Report.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return nil
}

func (t *ExternalRepositoryImpl) FindByName(name string) (*model.External, *helper.CustomError) {
	var err error
	external := model.External{}

	if err = t.Db.Model(model.External{}).First(&external, "name = ?", name).Error; err != nil {
		return nil, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}
	return &external, nil
}
