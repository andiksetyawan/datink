package swipe

import (
	"context"

	"datink/internal/models/auth"
	"datink/internal/models/swipe/dto"
	"datink/internal/repository"
	"datink/internal/resource"
)

type Service interface {
	Swiping(ctx context.Context, user *auth.UserJwtAuth, req dto.SwipingRequest) (dto.SwipingResponse, error)
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
