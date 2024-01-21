package user

import (
	"context"

	"datink/internal/models/auth"
	"datink/internal/models/user/dto"
	"datink/internal/models/user/entity"
)

func (s service) FindUserMatching(ctx context.Context, user *auth.UserJwtAuth, req dto.FindUserMatchingRequest) ([]entity.User, error) {
	return s.repository.User.FindUserMatching(ctx, s.resource.Database.GetMaster(), user.UserID, &req.Paging)
}
