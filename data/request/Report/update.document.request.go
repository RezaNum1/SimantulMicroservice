package report

type UpdateDokumenTemuanReportRequest struct {
	ID            string `validate:"required,min=1,max=200" json:"id"`
	Status        int    `json:"status"`
	DokumenTemuan string `json:"dokumenTemuan"`
}

type UpdateDokumenTindakLanjutReportRequest struct {
	ID                  string `validate:"required,min=1,max=200" json:"id"`
	Status              int    `json:"status"`
	DokumenTindakLanjut string `json:"dokumenTindakLanjut"`
}
