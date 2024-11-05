package report

import (
	"Expire/helper"
	"Expire/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReportRepositoryImpl struct {
	Db *gorm.DB
}

func NewReportRepositoryImpl(Db *gorm.DB) ReportRepository {
	return &ReportRepositoryImpl{Db: Db}
}

func (t *ReportRepositoryImpl) Create(report model.Report) *helper.CustomError {
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

func (t *ReportRepositoryImpl) GetReport(id string) (*model.Report, *helper.CustomError) {
	var report model.Report

	reportId, err := uuid.Parse(id)
	result := t.Db.Preload("Bank").Preload("Supervisor").Preload("Leader").First(&report, "id = ?", reportId)
	if result.Error != nil || err != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reports",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return &report, nil
}

func (t *ReportRepositoryImpl) GetAllReport() ([]model.Report, *helper.CustomError) {
	var reports []model.Report
	result := t.Db.Preload("Bank").Preload("Supervisor").Preload("Leader").Find(&reports)
	if result.Error != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reports",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return reports, nil
}

func (t *ReportRepositoryImpl) Update(report model.Report) *helper.CustomError {
	err := t.Db.Save(&report).Error

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

/*
Why when we want to get a spectifict data with 'where' clause, it cannot, while Find gives a result as we expected in gorm.db?
*/
