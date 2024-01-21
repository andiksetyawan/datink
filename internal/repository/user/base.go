package user

import (
	"context"

	"github.com/uptrace/bun"

	"datink/internal/models/common"
	"datink/internal/models/user/entity"
	"datink/internal/resource"
)

type Repository interface {
	Create(ctx context.Context, db bun.IDB, user *entity.User) error
	FindAll(ctx context.Context, db bun.IDB) ([]entity.User, error)
	GetByEmail(ctx context.Context, db bun.IDB, email string) (*entity.User, error)
	GetByID(ctx context.Context, db bun.IDB, id int64) (*entity.User, error)
	FindUserMatching(ctx context.Context, db bun.IDB, userID int64, opt *common.Paging) ([]entity.User, error)
	FindUserMatch(ctx context.Context, db bun.IDB, userID int64, opt *common.Paging) ([]entity.User, error)
}

type repository struct {
	resource resource.Resource
}

func NewRepository(resource resource.Resource) (Repository, error) {
	return &repository{
		resource: resource,
	}, nil
}
