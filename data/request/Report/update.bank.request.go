package report

import "time"

type UpdateBankReportRequest struct {
	ID                   string    `validate:"required,min=1,max=200" json:"id"`
	PoinTindakLanjut     string    `json:"poinTindakLanjut"`
	KomitmenTindakLanjut string    `json:"komitmenTindakLanjut"`
	WaktuPenyelesaian    time.Time `json:"waktuPenyelesaian"`
	Status               int       `json:"status"`
	DokumenTindakLanjut  string    `json:"dokumenTindakLanjut"`
}
