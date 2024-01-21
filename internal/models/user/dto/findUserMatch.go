package dto

import (
	"datink/internal/models/common"
	"datink/internal/models/user/entity"
)

type FindUserMatchRequest struct {
	common.Paging
}

type FindUserMatchResponse []entity.User
