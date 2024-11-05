package report

type FindReasonsRequest struct {
	ReportID string `validate:"required" json:"reportId"`
}
