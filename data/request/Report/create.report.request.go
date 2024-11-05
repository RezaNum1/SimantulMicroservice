package report

import "time"

type CreateReportRequest struct {
	Periode             string    `validate:"required,min=1,max=200" json:"periode"`
	JenisTemuan         string    `validate:"required,min=1,max=200" json:"jenisTemuan"`
	JudulTemuan         string    `validate:"required" json:"judulTemuan"`
	JenisPemeriksaan    string    `validate:"required" json:"jenisPemeriksaan"`
	JenisKantor         string    `validate:"required" json:"jenisKantor"`
	PoinTemuan          string    `validate:"required" json:"poinTemuan"`
	RincianTemuan       string    `validate:"required" json:"rincianTemuan"`
	RencanaTindakLanjut string    `validate:"required" json:"rencanaTindakLanjut"`
	DokumenTemuan       string    `json:"dokumenTemuan"`
	TanggalPemeriksaan  time.Time `validate:"required" json:"tanggalPemeriksaan"`
	TargetPenyelesaian  time.Time `validate:"required" json:"targetPenyelesaian"`
	BankID              string    `validate:"required" json:"bankID"`
	SupervisorID        string    `validate:"required" json:"supervisorID"`
	LeaderID            string    `validate:"required" json:"leaderID"`
}
