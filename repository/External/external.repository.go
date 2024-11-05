package external

import (
	"Expire/helper"
	"Expire/model"
)

type ExternalRepository interface {
	Create(external model.External) *helper.CustomError
	GetExternal(id string) (*model.External, *helper.CustomError)
	GetAllExternal() ([]model.External, *helper.CustomError)
	Update(report model.External) *helper.CustomError
	FindByName(name string) (*model.External, *helper.CustomError)
}
