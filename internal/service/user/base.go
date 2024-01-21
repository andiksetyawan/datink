package user

import (
	"context"

	"datink/internal/models/auth"
	"datink/internal/models/user/dto"
	"datink/internal/models/user/entity"
	"datink/internal/repository"
	"datink/internal/resource"
)

type Service interface {
	FindAll(ctx context.Context) ([]entity.User, error)
	SignUp(ctx context.Context, req dto.SignUpRequest) (dto.SignUpResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	FindUserMatching(ctx context.Context, user *auth.UserJwtAuth, req dto.FindUserMatchingRequest) ([]entity.User, error)
	FindUserMatch(ctx context.Context, user *auth.UserJwtAuth, req dto.FindUserMatchRequest) ([]entity.User, error)
}

type service struct {
	resource   resource.Resource
	repository repository.Repository
}

func NewService(resource resource.Resource, repository repository.Repository) (Service, error) {
	return &service{
		resource:   resource,
		repository: repository,
	}, nil
}
