package userPackage

import (
	"context"

	"datink/internal/models/auth"
	"datink/internal/models/userPackage/dto"
	"datink/internal/repository"
	"datink/internal/resource"
)

type Service interface {
	Purchase(ctx context.Context, user *auth.UserJwtAuth, req dto.UserPackagePurchaseRequest) (dto.UserPackagePurchaseResponse, error)
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
