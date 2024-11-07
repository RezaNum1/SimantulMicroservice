package leader

import (
	request "Expire/data/request/Leader"
	response "Expire/data/response/Leader"
	"Expire/helper"
)

type LeaderService interface {
	Create(report request.CreateLeaderRequest) *helper.CustomError
	GetLeader(id string) (*response.LeaderResponse, *helper.CustomError)
	GetAllLeader() ([]response.LeaderResponse, *helper.CustomError)
	Delete(id string) *helper.CustomError
}
