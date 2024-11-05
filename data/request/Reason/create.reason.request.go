package report

type CreateReasonRequest struct {
	Description string `validate:"required" json:"description"`
	ReportID    string `validate:"required" json:"reportId"`
}
