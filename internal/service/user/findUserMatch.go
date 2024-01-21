package user

import (
	"context"

	"datink/internal/models/auth"
	"datink/internal/models/user/dto"
	"datink/internal/models/user/entity"
)

func (s service) FindUserMatch(ctx context.Context, user *auth.UserJwtAuth, req dto.FindUserMatchRequest) ([]entity.User, error) {
	return s.repository.User.FindUserMatch(ctx, s.resource.Database.GetMaster(), user.UserID, &req.Paging)
}
