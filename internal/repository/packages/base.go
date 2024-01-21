package packages

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/package/entity"
	"datink/internal/resource"
)

type Repository interface {
	FindAll(ctx context.Context, db bun.IDB) ([]entity.Package, error)
	GetByID(ctx context.Context, db bun.IDB, id int64) (*entity.Package, error)
}

type repository struct {
	resource resource.Resource
}

func NewRepository(resource resource.Resource) (Repository, error) {
	return &repository{
		resource: resource,
	}, nil
}
