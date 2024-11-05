package leader

import (
	"Expire/helper"
	"Expire/model"
)

type LeaderRepository interface {
	Create(report model.Leader) *helper.CustomError
	GetLeader(id string) (*model.Leader, *helper.CustomError)
	GetAllLeader() ([]model.Leader, *helper.CustomError)
	FindLeaderById(leaderId string) (*model.Leader, *helper.CustomError)
	FindByName(name string) (*model.Leader, *helper.CustomError)
}
