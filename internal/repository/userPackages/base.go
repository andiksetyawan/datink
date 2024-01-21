package userPackages

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/userPackage/entity"
	"datink/internal/resource"
)

type Repository interface {
	FindAll(ctx context.Context, db bun.IDB) ([]entity.UserPackage, error)
	Create(ctx context.Context, db bun.IDB, user *entity.UserPackage) error
	GetActivePackage(ctx context.Context, db bun.IDB, userID int64) (*entity.UserPackage, error)
}

type repository struct {
	resource resource.Resource
}

func NewRepository(resource resource.Resource) (Repository, error) {
	return &repository{
		resource: resource,
	}, nil
}
