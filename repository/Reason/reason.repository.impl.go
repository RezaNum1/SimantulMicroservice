package reason

import (
	"Expire/helper"
	"Expire/model"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReasonRepositoryImpl struct {
	Db *gorm.DB
}

func NewReasonRepositoryImpl(Db *gorm.DB) ReasonRepository {
	return &ReasonRepositoryImpl{Db: Db}
}

func (t *ReasonRepositoryImpl) Create(report model.Reason) *helper.CustomError {
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

func (t *ReasonRepositoryImpl) GetReason(id string) (*model.Reason, *helper.CustomError) {
	var reason model.Reason

	reasonId, err := uuid.Parse(id)
	result := t.Db.First(&reason, "id = ?", reasonId)
	if result.Error != nil || err != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reason",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return &reason, nil
}

func (t *ReasonRepositoryImpl) GetAllReason() ([]model.Reason, *helper.CustomError) {
	var reasons []model.Reason
	result := t.Db.Find(&reasons)
	if result.Error != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reports",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return reasons, nil
}

func (t *ReasonRepositoryImpl) FindReasonsByReportID(reportId string) ([]model.Reason, *helper.CustomError) {
	var err error
	reason := []model.Reason{}

	if err = t.Db.Model(model.Reason{}).Find(&reason, "report_id = ?", strings.ToLower(reportId)).Error; err != nil {
		return nil, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}
	return reason, nil
}
