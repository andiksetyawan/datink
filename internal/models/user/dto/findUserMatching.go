package dto

import (
	"datink/internal/models/common"
	"datink/internal/models/user/entity"
)

type FindUserMatchingRequest struct {
	common.Paging
}

type FindUserMatchingResponse []entity.User
