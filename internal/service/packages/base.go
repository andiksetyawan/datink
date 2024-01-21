package packages

import (
	"context"

	"datink/internal/models/package/entity"
	"datink/internal/repository"
	"datink/internal/resource"
)

type Service interface {
	FindAll(ctx context.Context) ([]entity.Package, error)
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
