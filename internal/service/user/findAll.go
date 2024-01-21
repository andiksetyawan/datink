package user

import (
	"context"

	"datink/internal/models/user/entity"
)

func (s service) FindAll(ctx context.Context) ([]entity.User, error) {
	customers, err := s.repository.User.FindAll(ctx, s.resource.Database.GetMaster())
	if err != nil {
		return nil, err
	}

	return customers, nil
}
