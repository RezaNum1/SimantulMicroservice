package leader

type CreateLeaderRequest struct {
	Name    string `validate:"required,min=1,max=200" json:"name"`
	Jabatan string `validate:"required,min=1,max=200" json:"jabatan"`
	NIP     string `validate:"required,min=1,max=200" json:"nip"`
	Phone   string `validate:"required,min=1,max=200" json:"phone"`
}
