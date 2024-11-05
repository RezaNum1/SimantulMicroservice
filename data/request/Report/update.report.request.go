package report

import "time"

type UpdateReportRequest struct {
	ID                  string    `validate:"required,min=1,max=200" json:"id"`
	Periode             string    `validate:"required,min=1,max=200" json:"periode"`
	JenisTemuan         string    `validate:"required,min=1,max=200" json:"jenisTemuan"`
	JudulTemuan         string    `validate:"required" json:"judulTemuan"`
	JenisPemeriksaan    string    `validate:"required" json:"jenisPemeriksaan"`
	JenisKantor         string    `validate:"required" json:"jenisKantor"`
	PoinTemuan          string    `validate:"required" json:"poinTemuan"`
	RincianTemuan       string    `validate:"required" json:"rincianTemuan"`
	RencanaTindakLanjut string    `json:"rencanaTindakLanjut"`
	DokumenTemuan       string    `json:"dokumenTemuan"`
	TanggalPemeriksaan  time.Time `validate:"required" json:"tanggalPemeriksaan"`
	TargetPenyelesaian  time.Time `validate:"required" json:"targetPenyelesaian"`
	BankID              string    `validate:"required" json:"bankID"`
	SupervisorID        string    `validate:"required" json:"supervisorID"`
	Status              int       `json:"status"`
	PrevStatus          int       `json:"prevStatus"`
}
