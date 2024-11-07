package report

import (
	request "Expire/data/request/Report"
	response "Expire/data/response/Report"
	"Expire/helper"
	"Expire/model"
	bankRepository "Expire/repository/Bank"
	externalRepository "Expire/repository/External"
	leaderRepository "Expire/repository/Leader"
	reasonRepository "Expire/repository/Reason"
	repository "Expire/repository/Report"
	supervisorRepository "Expire/repository/Supervisor"

	"unicode/utf8"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ReportServiceImpl struct {
	ReportRepository     repository.ReportRepository
	SupervisorRepository supervisorRepository.SupervisorRepository
	LeaderRepository     leaderRepository.LeaderRepository
	BankRepository       bankRepository.BankRepository
	ReasonRepository     reasonRepository.ReasonRepository
	ExternalRepository   externalRepository.ExternalRepository
	Validate             *validator.Validate
}

func NewReportServiceImpl(
	reportRepository repository.ReportRepository,
	supervisorRepository supervisorRepository.SupervisorRepository,
	leaderRepository leaderRepository.LeaderRepository,
	bankRepository bankRepository.BankRepository,
	reasonRepository reasonRepository.ReasonRepository,
	externalRepository externalRepository.ExternalRepository,
	validate *validator.Validate) ReportService {
	return &ReportServiceImpl{
		ReportRepository:     reportRepository,
		SupervisorRepository: supervisorRepository,
		LeaderRepository:     leaderRepository,
		BankRepository:       bankRepository,
		ReasonRepository:     reasonRepository,
		ExternalRepository:   externalRepository,
		Validate:             validate,
	}
}

func (t ReportServiceImpl) Create(report request.CreateReportRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(report)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	// Get Supervisor ID
	supervisorResult, err := t.SupervisorRepository.FindSupervisorById(report.SupervisorID)
	if err != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	// Get Leader ID
	leaderResult, err := t.LeaderRepository.FindLeaderById(report.LeaderID)
	if err != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	// Get Bank ID
	bankResult, err := t.BankRepository.FindBankById(report.BankID)
	if err != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	// Get External ID
	// externalResult, err := t.ExternalRepository.GetExternal(report.ExternalID)
	// if err != nil {
	// 	fileName, atLine := helper.GetFileAndLine(errStructure)
	// 	return &helper.CustomError{
	// 		Code:     400,
	// 		Message:  "Invalid Request Structure.",
	// 		FileName: fileName,
	// 		AtLine:   atLine,
	// 	}
	// }

	reportModel := model.Report{
		Periode:             report.Periode,
		JenisTemuan:         report.JenisTemuan,
		JudulTemuan:         report.JudulTemuan,
		JenisPemeriksaan:    report.JenisPemeriksaan,
		JenisKantor:         report.JenisKantor,
		PoinTemuan:          report.PoinTemuan,
		RincianTemuan:       report.RincianTemuan,
		RencanaTindakLanjut: report.RencanaTindakLanjut,
		DokumenTemuan:       report.DokumenTemuan,
		TanggalPemeriksaan:  report.TanggalPemeriksaan,
		TargetPenyelesaian:  report.TargetPenyelesaian,
		SupervisorID:        *supervisorResult.ID,
		BankID:              *bankResult.ID,
		LeaderID:            *leaderResult.ID,
	}
	createError := t.ReportRepository.Create(reportModel)

	if createError != nil {
		println("🐶 11")
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

func (t ReportServiceImpl) GetReport(id string) (*response.ReportResponse, *helper.CustomError) {
	result, fetchError := t.ReportRepository.GetReport(id)

	if fetchError != nil {
		return nil, fetchError
	} else {
		response := response.ReportResponse{
			Id:                   (*uuid.UUID)(result.ID),
			Periode:              result.Periode,
			JenisTemuan:          result.JenisTemuan,
			JudulTemuan:          result.JudulTemuan,
			JenisPemeriksaan:     result.JenisPemeriksaan,
			JenisKantor:          result.JenisKantor,
			PoinTemuan:           result.PoinTemuan,
			RincianTemuan:        result.RincianTemuan,
			RencanaTindakLanjut:  result.RencanaTindakLanjut,
			DokumenTemuan:        result.DokumenTemuan,
			TanggalPemeriksaan:   result.TanggalPemeriksaan,
			TargetPenyelesaian:   result.TargetPenyelesaian,
			PoinTindakLanjut:     result.PoinTindakLanjut,
			KomitmenTindakLanjut: result.KomitmenTindakLanjut,
			DokumenTindakLanjut:  result.DokumenTindakLanjut,
			WaktuPenyelesaian:    result.WaktuPenyelesaian,
			Status:               result.Status,
			PrevStatus:           result.PrevStatus,
			SupervisorID:         result.SupervisorID.String(),
			BankID:               result.BankID.String(),
			Bank:                 result.Bank,
			Supervisor:           result.Supervisor,
			LeaderID:             result.LeaderID.String(),
			Leader:               result.Leader,
		}

		return &response, nil
	}
}

func (t ReportServiceImpl) GetAllReport() ([]response.ReportResponse, *helper.CustomError) {
	result, fetchError := t.ReportRepository.GetAllReport()

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapReportsToReportResponse(result), nil
	}
}

func (t ReportServiceImpl) GetAllSupervisorReports(id string) ([]response.ReportResponse, *helper.CustomError) {
	result, fetchError := t.ReportRepository.GetAllSupervisorReports(id)

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapReportsToReportResponse(result), nil
	}
}

func (t ReportServiceImpl) GetAllLeaderReports(id string) ([]response.ReportResponse, *helper.CustomError) {
	result, fetchError := t.ReportRepository.GetAllLeaderReports(id)

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapReportsToReportResponse(result), nil
	}
}

func (t ReportServiceImpl) GetAllBankReports(id string) ([]response.ReportResponse, *helper.CustomError) {
	externals, _ := t.ExternalRepository.GetExternal(id)
	println(externals.Bank.ID.String())
	result, fetchError := t.ReportRepository.GetAllBankReports(externals.Bank.ID.String())

	if fetchError != nil {
		return nil, fetchError
	} else {
		return t.mapReportsToReportResponse(result), nil
	}
}

func (t ReportServiceImpl) mapReportsToReportResponse(reports []model.Report) []response.ReportResponse {
	responseReports := make([]response.ReportResponse, len(reports))
	for i, report := range reports {
		responseReports[i] = t.convertReportToReportResponse(report)
	}
	return responseReports
}

func (t ReportServiceImpl) convertReportToReportResponse(report model.Report) response.ReportResponse {
	// Perform necessary conversion logic here, potentially selecting specific fields
	responseReport := response.ReportResponse{
		Id:                   (*uuid.UUID)(report.ID),
		Periode:              report.Periode,
		JenisTemuan:          report.JenisTemuan,
		JudulTemuan:          report.JudulTemuan,
		JenisPemeriksaan:     report.JenisPemeriksaan,
		JenisKantor:          report.JenisKantor,
		PoinTemuan:           report.PoinTemuan,
		RincianTemuan:        report.RincianTemuan,
		RencanaTindakLanjut:  report.RencanaTindakLanjut,
		DokumenTemuan:        report.DokumenTemuan,
		TanggalPemeriksaan:   report.TanggalPemeriksaan,
		TargetPenyelesaian:   report.TargetPenyelesaian,
		PoinTindakLanjut:     report.PoinTindakLanjut,
		KomitmenTindakLanjut: report.KomitmenTindakLanjut,
		DokumenTindakLanjut:  report.DokumenTindakLanjut,
		WaktuPenyelesaian:    report.WaktuPenyelesaian,
		Status:               report.Status,
		PrevStatus:           report.PrevStatus,
		SupervisorID:         report.SupervisorID.String(),
		BankID:               report.BankID.String(),
		Bank:                 report.Bank,
		Supervisor:           report.Supervisor,
		Leader:               report.Leader,
	}
	return responseReport
}

func (t ReportServiceImpl) Update(report request.UpdateReportRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(report)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	// Get Supervisor ID
	supervisorResult, err := t.SupervisorRepository.FindSupervisorById(report.SupervisorID)

	if err != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	// Get Bank ID
	bankResult, err := t.BankRepository.FindBankById(report.BankID)

	if err != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	data, fetchErr := t.ReportRepository.GetReport(report.ID)
	data.Periode = report.Periode
	data.JenisTemuan = report.JenisTemuan
	data.JudulTemuan = report.JudulTemuan
	data.JenisPemeriksaan = report.JenisPemeriksaan
	data.JenisKantor = report.JenisKantor
	data.PoinTemuan = report.PoinTemuan
	data.RincianTemuan = report.RincianTemuan
	data.RencanaTindakLanjut = report.RencanaTindakLanjut
	data.TanggalPemeriksaan = report.TanggalPemeriksaan
	data.TargetPenyelesaian = report.TargetPenyelesaian
	data.Status = report.Status
	data.PrevStatus = report.PrevStatus
	data.BankID = *bankResult.ID
	data.SupervisorID = *supervisorResult.ID

	saveErr := t.ReportRepository.Update(*data)

	if fetchErr != nil || saveErr != nil {
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

func (t ReportServiceImpl) UpdateStatus(report request.UpdateStatusReportRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(report)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	data, fetchErr := t.ReportRepository.GetReport(report.ID)
	// Add Resons
	if utf8.RuneCountInString(report.Alasan) > 0 && report.Alasan != "" {
		errReason := t.ReasonRepository.Create(model.Reason{Description: report.Alasan, RejectedStep: data.Status, ReportID: data.ID.String()})
		if errReason != nil {
			return errReason
		}
	}

	data.PrevStatus = data.Status
	data.Status = report.Status

	saveErr := t.ReportRepository.Update(*data)

	if fetchErr != nil || saveErr != nil {
		println("🐶 2")
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

func (t ReportServiceImpl) UpdateBankReport(report request.UpdateBankReportRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(report)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	data, fetchErr := t.ReportRepository.GetReport(report.ID)
	data.Status = report.Status
	data.PoinTindakLanjut = report.PoinTindakLanjut
	data.KomitmenTindakLanjut = report.KomitmenTindakLanjut
	data.DokumenTindakLanjut = report.DokumenTindakLanjut
	data.WaktuPenyelesaian = report.WaktuPenyelesaian
	data.PrevStatus = (report.Status - 1)

	saveErr := t.ReportRepository.Update(*data)

	if fetchErr != nil || saveErr != nil {
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

func (t ReportServiceImpl) UpdateDokumenTemuanStatus(report request.UpdateDokumenTemuanReportRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(report)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	data, fetchErr := t.ReportRepository.GetReport(report.ID)
	data.PrevStatus = data.Status
	data.Status = report.Status
	data.DokumenTemuan = report.DokumenTemuan

	saveErr := t.ReportRepository.Update(*data)

	if fetchErr != nil || saveErr != nil {
		println("🐶 2")
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

func (t ReportServiceImpl) UpdateDokumenTindakLanjutStatus(report request.UpdateDokumenTindakLanjutReportRequest) *helper.CustomError {
	errStructure := t.Validate.Struct(report)
	if errStructure != nil {
		fileName, atLine := helper.GetFileAndLine(errStructure)
		return &helper.CustomError{
			Code:     400,
			Message:  "Invalid Request Structure",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	data, fetchErr := t.ReportRepository.GetReport(report.ID)
	data.PrevStatus = data.Status
	data.Status = report.Status
	data.DokumenTindakLanjut = report.DokumenTindakLanjut

	saveErr := t.ReportRepository.Update(*data)

	if fetchErr != nil || saveErr != nil {
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
