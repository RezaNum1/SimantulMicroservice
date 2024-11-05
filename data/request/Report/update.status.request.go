package report

type UpdateStatusReportRequest struct {
	ID     string `validate:"required,min=1,max=200" json:"id"`
	Status int    `json:"status"`
	Alasan string `json:"alasan"`
}
