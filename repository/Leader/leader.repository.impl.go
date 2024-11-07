package leader

import (
	"Expire/helper"
	"Expire/model"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LeaderRepositoryImpl struct {
	Db *gorm.DB
}

func NewLeaderRepositoryImpl(Db *gorm.DB) LeaderRepository {
	return &LeaderRepositoryImpl{Db: Db}
}

func (t *LeaderRepositoryImpl) Create(report model.Leader) *helper.CustomError {
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

func (t *LeaderRepositoryImpl) GetLeader(id string) (*model.Leader, *helper.CustomError) {
	var leader model.Leader

	leaderId, err := uuid.Parse(id)
	result := t.Db.First(&leader, "id = ?", leaderId)
	if result.Error != nil || err != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Leader",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return &leader, nil
}

func (t *LeaderRepositoryImpl) GetAllLeader() ([]model.Leader, *helper.CustomError) {
	var leaders []model.Leader
	result := t.Db.Find(&leaders)
	if result.Error != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reports",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return leaders, nil
}

func (t *LeaderRepositoryImpl) FindLeaderById(leaderId string) (*model.Leader, *helper.CustomError) {
	var err error
	leader := model.Leader{}

	if err = t.Db.Model(model.Leader{}).First(&leader, "id = ?", strings.ToLower(leaderId)).Error; err != nil {
		return nil, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}
	return &leader, nil
}

func (t *LeaderRepositoryImpl) FindByName(name string) (*model.Leader, *helper.CustomError) {
	var err error
	leader := model.Leader{}

	if err = t.Db.Model(model.Leader{}).First(&leader, "name = ?", name).Error; err != nil {
		return nil, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}
	return &leader, nil
}

func (t *LeaderRepositoryImpl) Delete(id string) *helper.CustomError {
	userId, err := uuid.Parse(id)
	result := t.Db.Unscoped().Delete(&model.Leader{}, userId)

	if result.Error != nil || err != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reports",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return nil
}
