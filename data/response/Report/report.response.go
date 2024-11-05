package report

import (
	"Expire/model"
	"time"

	"github.com/google/uuid"
)

type ReportResponse struct {
	Id                   *uuid.UUID       `json:"id"`
	Periode              string           `json:"periode"`
	JenisTemuan          string           `json:"jenisTemuan"`
	JudulTemuan          string           `json:"judulTemuan"`
	JenisPemeriksaan     string           `json:"jenisPemeriksaan"`
	JenisKantor          string           `json:"jenisKantor"`
	PoinTemuan           string           `json:"poinTemuan"`
	RincianTemuan        string           `json:"rincianTemuan"`
	RencanaTindakLanjut  string           `json:"rencanaTindakLanjut"`
	DokumenTemuan        string           `json:"dokumenTemuan"`
	TanggalPemeriksaan   time.Time        `json:"tanggalPemeriksaan"`
	TargetPenyelesaian   time.Time        `json:"targetPenyelesaian"`
	PoinTindakLanjut     string           `json:"poinTindakLanjut"`
	KomitmenTindakLanjut string           `json:"komitmenTindakLanjut"`
	DokumenTindakLanjut  string           `json:"dokumenTindakLanjut"`
	WaktuPenyelesaian    time.Time        `json:"waktuPenyelesaian"`
	Status               int              `json:"status"`
	PrevStatus           int              `json:"prevStatus"`
	BankID               string           `json:"bankID"`
	SupervisorID         string           `json:"supervisorID"`
	LeaderID             string           `json:"leaderID"`
	ExternalID           string           `json:"externalID"`
	Bank                 model.Bank       `json:"bank"`
	Supervisor           model.Supervisor `json:"supervisor"`
	Leader               model.Leader     `json:"leader"`
	External             model.External   `json:"external"`
}
