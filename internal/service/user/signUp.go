package user

import (
	"context"

	"datink/internal/models/user/dto"
	"datink/internal/models/user/entity"
	"datink/pkg/util"
)

func (s service) SignUp(ctx context.Context, req dto.SignUpRequest) (dto.SignUpResponse, error) {
	user := entity.User{
		Name:      req.Name,
		Email:     req.Email,
		Birthdate: req.Birthdate,
		Pictures:  nil,
	}

	password, err := util.HashPassword(req.Password)
	if err != nil {
		return dto.SignUpResponse{}, err
	}

	user.Password = password

	err = s.repository.User.Create(ctx, s.resource.Database.GetMaster(), &user)
	if err != nil {
		return dto.SignUpResponse{}, err
	}

	return dto.SignUpResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
